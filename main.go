package main

import (
	"Reactive-Welfare-Housing-System/src/distributor"
	"Reactive-Welfare-Housing-System/src/manager"
	"Reactive-Welfare-Housing-System/src/verifier"
	"encoding/json"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/persistence"
	"strconv"

	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/storage"
	"log"
	"net/http"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/mailbox"
)

type Provider struct {
	providerState persistence.ProviderState
}

func NewProvider(snapshotInterval int) *Provider {
	return &Provider{
		providerState: persistence.NewInMemoryProvider(snapshotInterval),
	}
}

func (p *Provider) InitState(actorName string, eventNum, eventIndexAfterSnapshot int) {
	for i := 0; i < eventNum; i++ {
		p.providerState.PersistEvent(
			actorName,
			i,
			&Message{protoMsg: protoMsg{state: "state" + strconv.Itoa(i)}},
		)
	}
	p.providerState.PersistSnapshot(
		actorName,
		eventIndexAfterSnapshot,
		&Snapshot{protoMsg: protoMsg{state: "state" + strconv.Itoa(eventIndexAfterSnapshot-1)}},
	)
}

func (p *Provider) GetState() persistence.ProviderState {
	return p.providerState
}

func (p *protoMsg) Reset()         {}
func (p *protoMsg) String() string { return p.state }
func (p *protoMsg) ProtoMessage()  {}

type protoMsg struct{ state string }
type Message struct{ protoMsg }
type Snapshot struct{ protoMsg }

type Env struct {
	db             storage.HouseSystem
	distributorPID *actor.PID
	managerPID     *actor.PID
	verifierPID    *actor.PID
}

var system *actor.ActorSystem

func main() {
	db, err := storage.InitDB()
	if err != nil {
		log.Panic(err)
	}

	system = actor.NewActorSystem()
	distributorProvider := NewProvider(3)
	distributorProvider.InitState("distributor", 4, 3)
	distributorPID := system.Root.Spawn(actor.PropsFromProducer(distributor.NewDistributorActor(db)).WithMailbox(mailbox.Unbounded()).WithReceiverMiddleware(persistence.Using(distributorProvider)))

	managerProvider := NewProvider(3)
	managerProvider.InitState("manager", 4, 3)
	managerPID := system.Root.Spawn(actor.PropsFromProducer(manager.NewManagerActor(db)).WithMailbox(mailbox.Unbounded()).WithReceiverMiddleware(persistence.Using(managerProvider)))

	verifierProvider := NewProvider(3)
	verifierProvider.InitState("verifier", 4, 3)
	veriferPID := system.Root.Spawn(actor.PropsFromProducer(verifier.NewVerifierActor()).WithMailbox(mailbox.Unbounded()).WithReceiverMiddleware(persistence.Using(verifierProvider)))
	fmt.Println(distributorPID, managerPID, veriferPID)
	system.Root.Send(distributorPID, &sharedMessages.ManagerConnect{
		Sender: managerPID,
	})
	system.Root.Send(veriferPID, &sharedMessages.DistributorConnect{
		Sender: distributorPID,
	})
	system.Root.Send(managerPID, &sharedMessages.VerifierConnect{
		Sender: veriferPID,
	})
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
		system.Root.Request(env.verifierPID, &sharedMessages.NewRequests{Requests: requests})
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
		system.Root.Request(env.verifierPID, &sharedMessages.NewCheckOuts{CheckOuts: checkouts})
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
		system.Root.Request(env.managerPID, &sharedMessages.ExaminationList{HouseID: checkids})
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
		system.Root.Request(env.managerPID, &sharedMessages.NewHouses{Houses: houses})
		// distributor.
		// Do something with POST URL
		// messages.
	}
}
