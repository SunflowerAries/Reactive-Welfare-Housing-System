package main

import (
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
	fmt.Println(distributorPID)
	http.HandleFunc("/houses", env.housesIndex)
	http.ListenAndServe(":3000", nil)
}

func (env *Env) housesIndex(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		// fmt.Fprintf(w, "Not Implement\n")
		// Do something with GET URL
	case "POST":
		// distributor.
		// Do something with POST URL
	}

}
