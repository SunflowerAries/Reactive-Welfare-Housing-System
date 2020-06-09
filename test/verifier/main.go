package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ChaokunChang/protoactor-go/actor"
	"github.com/ChaokunChang/protoactor-go/mailbox"
	"github.com/ChaokunChang/protoactor-go/remote"

	"Reactive-Welfare-Housing-System/src/shared"
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
	default:
		fmt.Printf("Unexpected message for VerifierSupervisor: %+v\n", msg)
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
