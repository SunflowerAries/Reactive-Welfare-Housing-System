package main

import (
	"encoding/json"
	"housingSystem/src/distributor"
	"housingSystem/src/manager"
	"housingSystem/src/messages"
	"housingSystem/src/storage"
	"log"
	"net/http"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/mailbox"
)

type Env struct {
	db             storage.HouseSystem
	distributorPID *actor.PID
	managerPID     *actor.PID
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
	env := &Env{db, distributorPID, managerPID}
	http.HandleFunc("/houses", env.housesIndex)
	http.ListenAndServe(":3000", nil)
}

func (env *Env) housesIndex(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		// fmt.Fprintf(w, "Not Implement\n")
		// Do something with GET URL
	case "POST":
		decoder := json.NewDecoder(req.Body)
		decoder.DisallowUnknownFields()
		house := []storage.House{}
		err := decoder.Decode(&house)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if decoder.More() {
			http.Error(w, "extraneous data after JSON object", http.StatusBadRequest)
		}
		log.Printf("%+v\n", house)
		houses := make([]*messages.House, len(house))
		for i := range house {
			houses[i] = &messages.House{Level: int32(house[i].Level), Age: int32(house[i].Age), Area: int32(house[i].Area)}
		}
		rootContext.Request(env.managerPID, &messages.Houses{Houses: houses})
		// distributor.
		// Do something with POST URL
	}

}
