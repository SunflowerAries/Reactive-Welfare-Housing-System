package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ChaokunChang/protoactor-go/actor"
	"github.com/ChaokunChang/protoactor-go/mailbox"
	"github.com/ChaokunChang/protoactor-go/remote"

	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/shared"
	"Reactive-Welfare-Housing-System/src/storage"
	"Reactive-Welfare-Housing-System/src/verifier"
)

type verifierSupervisor struct {
	verifier *actor.PID
}

func newVerifierSupervisor() actor.Actor {
	return &verifierSupervisor{}
}

var rootContext = actor.EmptyRootContext

func (state *verifierSupervisor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("VerifierSupervisor Started, PID:", ctx.Self())
		state.verifier, _ = rootContext.SpawnNamed(
			actor.PropsFromProducer(verifier.NewVerifierActor()).WithMailbox(mailbox.Unbounded()),
			"Verifier")
		fmt.Println("Verifier Started, PID:", state.verifier)

		http.HandleFunc("/users", state.usersIndex)
		http.HandleFunc("/users/checkout", state.userscheckoutIndex)
		http.ListenAndServe(":4000", nil)
	default:
		fmt.Printf("Unexpected message for VerifierSupervisor: %+v\n", msg)
	}
}

func (state *verifierSupervisor) usersIndex(w http.ResponseWriter, req *http.Request) {
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
		rootContext.Request(state.verifier, &sharedMessages.NewRequests{Requests: requests})
	}
}

func (state *verifierSupervisor) userscheckoutIndex(w http.ResponseWriter, req *http.Request) {
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
		rootContext.Request(state.verifier, &sharedMessages.NewCheckOuts{CheckOuts: checkouts})
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	remote.Start("127.0.0.1:9001")

	decider := func(reason interface{}) actor.Directive {
		fmt.Println("handling failure for child")
		return actor.StopDirective
	}
	supervisor := actor.NewOneForOneStrategy(10, 1000, decider)
	props := actor.
		PropsFromProducer(newVerifierSupervisor).
		WithSupervisor(supervisor).
		WithMailbox(mailbox.Unbounded())

	pid, _ := rootContext.SpawnNamed(props, "VerifierSupervisor")
	shared.Use(pid)
	text, _ := reader.ReadString('\n')
	shared.Use((text))
}
