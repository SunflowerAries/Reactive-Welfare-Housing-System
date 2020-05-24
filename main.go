package main

import (
	"encoding/json"
	"fmt"
	"housingSystem/src/distributor"
	"housingSystem/src/storage"
	"log"
	"net/http"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/mailbox"
)

type Env struct {
	db storage.HouseSystem
}

func main() {
	db, err := storage.InitDB()
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db}
	rootContext := actor.EmptyRootContext
	distributorPID := rootContext.Spawn(actor.PropsFromProducer(distributor.NewDistributorActor(env.db)).WithMailbox(mailbox.Unbounded()))
	// managerPID := rootContext.Spawn(actor.PropsFromProducer(manager.NewManagerActor(env.db)).WithMailbox(mailbox.Unbounded()))
	fmt.Println(distributorPID)
	fmt.Println("Hello world")
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
		// distributor.
		// Do something with POST URL
	}

}
