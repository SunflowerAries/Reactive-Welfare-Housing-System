package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/json"

	"github.com/ChaokunChang/protoactor-go/actor"
	"github.com/ChaokunChang/protoactor-go/mailbox"
	"github.com/ChaokunChang/protoactor-go/remote"

	"Reactive-Welfare-Housing-System/src/manager"
	"Reactive-Welfare-Housing-System/src/shared"
	"Reactive-Welfare-Housing-System/src/storage"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
)

type managerSupervisor struct {
	manager *actor.PID
}

func newManagerSupervisor() actor.Actor {
	return &managerSupervisor{}
}

var rootContext = actor.EmptyRootContext

func (state *managerSupervisor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("ManagerSupervisor Started, PID:", ctx.Self())
		db, err := storage.InitDB()
		if err != nil {
			log.Panic(err)
		}
		state.manager, _ = rootContext.SpawnNamed(
			actor.PropsFromProducer(manager.NewManagerActor(db)).WithMailbox(mailbox.Unbounded()),
			"Manager")
		fmt.Println("Manager Started, PID:", state.manager)
		http.HandleFunc("/houses", state.housesIndex)
		http.HandleFunc("/houses/check", state.housescheckIndex)
		http.ListenAndServe(":3000", nil)
	default:
		fmt.Printf("Unexpected message for ManagerSupervisor: %+v\n", msg)
	}
}


func (state *managerSupervisor) housescheckIndex(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		decoder := json.NewDecoder(req.Body)
		decoder.DisallowUnknownFields()
		var checkids []int32
		err := decoder.Decode(&checkids)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if decoder.More() {
			http.Error(w, "extraneous data after JSON object", http.StatusBadRequest)
		}
		// https://stackoverflow.com/questions/55381710/converting-internal-go-struct-array-to-protobuf-generated-pointer-array
		rootContext.Request(state.manager, &sharedMessages.ExaminationList{HouseID: checkids})
		// distributor.
		// Do something with POST URL
		// messages.
	}
}

func (state *managerSupervisor) housesIndex(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// fmt.Fprintf(w, "Not Implement\n")
		// Do something with GET URL
	case "POST":
		fmt.Print("In house")
		// https://stackoverflow.com/questions/15672556/handling-json-post-request-in-go
		decoder := json.NewDecoder(req.Body)
		decoder.DisallowUnknownFields()
		newhouse := []storage.House{}
		err := decoder.Decode(&newhouse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if decoder.More() {
			http.Error(w, "extraneous data after JSON object", http.StatusBadRequest)
		}
		// https://stackoverflow.com/questions/55381710/converting-internal-go-struct-array-to-protobuf-generated-pointer-array
		houses := make([]*sharedMessages.NewHouse, len(newhouse))
		for i, house := range newhouse {
			houses[i] = &sharedMessages.NewHouse{Level: house.Level, Age: house.Age, Area: house.Area}
		}
		rootContext.Request(state.manager, &sharedMessages.NewHouses{Houses: houses})
		// distributor.
		// Do something with POST URL
		// messages.
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	remote.Start("127.0.0.1:9003")
	decider := func(reason interface{}) actor.Directive {
		fmt.Println("handling failure for child")
		return actor.StopDirective
	}
	supervisor := actor.NewOneForOneStrategy(10, 1000, decider)
	props := actor.
		PropsFromProducer(newManagerSupervisor).
		WithSupervisor(supervisor).
		WithMailbox(mailbox.Unbounded())

	pid, _ := rootContext.SpawnNamed(props, "ManagerSupervisor")
	shared.Use(pid)
	text, _ := reader.ReadString('\n')
	shared.Use((text))
}
