package main

import (
	"Reactive-Welfare-Housing-System/src/distributor"
	"Reactive-Welfare-Housing-System/src/manager"
	"Reactive-Welfare-Housing-System/src/verifier"
	"encoding/json"
	"fmt"

	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/storage"
	"log"
	"net/http"

	"github.com/ChaokunChang/protoactor-go/actor"
	"github.com/ChaokunChang/protoactor-go/mailbox"
)

// Env that stores the meta for this project
type Env struct {
	db             storage.HouseSystem
	distributorPID *actor.PID
	managerPID     *actor.PID
	verifierPID    *actor.PID
}

var rootContext *actor.RootContext

func main() {
	db, err := storage.InitDB()
	if err != nil {
		log.Panic(err)
	}

	rootContext = actor.EmptyRootContext
	distributorPID := rootContext.Spawn(actor.PropsFromProducer(distributor.NewDistributorActor(db)).WithMailbox(mailbox.Unbounded()))
	managerPID := rootContext.Spawn(actor.PropsFromProducer(manager.NewManagerActor(db)).WithMailbox(mailbox.Unbounded()))
	veriferPID := rootContext.Spawn(actor.PropsFromProducer(verifier.NewVerifierActor()).WithMailbox(mailbox.Unbounded()))
	fmt.Println(distributorPID, managerPID, veriferPID)
	rootContext.Send(distributorPID, &sharedMessages.ManagerConnect{
		Sender: managerPID,
	})
	rootContext.Send(veriferPID, &sharedMessages.DistributorConnect{
		Sender: distributorPID,
	})
	rootContext.Send(managerPID, &sharedMessages.VerifierConnect{
		Sender: veriferPID,
	})
	// rootContext.Send(managerPID, &sharedMessages.Connect{
	// 	Sender: distributorPID,
	// })
	env := &Env{db, distributorPID, managerPID, veriferPID}
	http.HandleFunc("/users/checkout", env.userscheckoutIndex)
	http.HandleFunc("/users", env.usersIndex)
	http.HandleFunc("/houses", env.housesIndex)
	http.HandleFunc("/houses/check", env.housescheckIndex)
	http.ListenAndServe(":3000", nil)
}

func (env *Env) usersIndex(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		decoder := json.NewDecoder(req.Body)
		decoder.DisallowUnknownFields()
		var newrequest []storage.FamilyCheckOut
		err := decoder.Decode(&newrequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if decoder.More() {
			http.Error(w, "extraneous data after JSON object", http.StatusBadRequest)
		}
		requests := make([]*sharedMessages.NewRequest, len(newrequest))
		for i, request := range newrequest {
			requests[i] = &sharedMessages.NewRequest{FamilyID: request.FamilyID, Level: request.Level}
		}
		rootContext.Request(env.verifierPID, &sharedMessages.NewRequests{Requests: requests})
	}
}

func (env *Env) userscheckoutIndex(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		decoder := json.NewDecoder(req.Body)
		decoder.DisallowUnknownFields()
		var newcheckout []storage.FamilyCheckOut
		err := decoder.Decode(&newcheckout)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if decoder.More() {
			http.Error(w, "extraneous data after JSON object", http.StatusBadRequest)
		}
		checkouts := make([]*sharedMessages.NewCheckOut, len(newcheckout))
		for i, checkout := range newcheckout {
			checkouts[i] = &sharedMessages.NewCheckOut{FamilyID: checkout.FamilyID, Level: checkout.Level}
		}
		rootContext.Request(env.verifierPID, &sharedMessages.NewCheckOuts{CheckOuts: checkouts})
	}
}

func (env *Env) housescheckIndex(w http.ResponseWriter, req *http.Request) {
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
		rootContext.Request(env.managerPID, &sharedMessages.ExaminationList{HouseID: checkids})
		// distributor.
		// Do something with POST URL
		// messages.
	}
}

func (env *Env) housesIndex(w http.ResponseWriter, req *http.Request) {
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
		rootContext.Request(env.managerPID, &sharedMessages.NewHouses{Houses: houses})
		// distributor.
		// Do something with POST URL
		// messages.
	}
}
