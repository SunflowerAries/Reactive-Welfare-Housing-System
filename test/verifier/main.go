package main

import (
	"fmt"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"

	"housingSystem/src/verifier"
	"housingSystem/src/shared"
)

type parentActor struct{}

func newParentActor() actor.Actor {
	return &parentActor{}
}

func (state *parentActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("VerifierParent Starting, initialize actor here, PID:", ctx.Self())
		shared.Use(msg)
		var rootContext = actor.EmptyRootContext
		props := actor.PropsFromProducer(func() actor.Actor { return &verifier.Actor{} })
		rootContext.SpawnNamed(props, "Verifier")
	}
}

func main() {
	decider := func(reason interface{}) actor.Directive {
		fmt.Println("handling failure for child")
		return actor.StopDirective
	}
	supervisor := actor.NewOneForOneStrategy(10, 1000, decider)
	rootContext := actor.EmptyRootContext
	props := actor.
		PropsFromProducer(newParentActor).
		WithSupervisor(supervisor)

	pid := rootContext.Spawn(props)
	shared.Use(pid)
	console.ReadLine()
}
