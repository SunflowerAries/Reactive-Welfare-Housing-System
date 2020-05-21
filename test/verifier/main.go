package main

import (
	"fmt"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"

	"Reactive-Welfare-Housing-System/src/shared"
	"Reactive-Welfare-Housing-System/src/verifier"
)

type parentActor struct{}

func newParentActor() actor.Actor {
	return &parentActor{}
}

func (state *parentActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		// timeout := 5 * time.Second
		fmt.Println("VerifierParent Starting, initialize actor here, PID:", ctx.Self())
		shared.Use(msg)
		var rootContext = actor.EmptyRootContext
		props := actor.PropsFromProducer(func() actor.Actor { return &verifier.Actor{} })
		rootContext.SpawnNamed(props, "Verifier")

		// pidResp, _ := remote.SpawnNamed("127.0.0.1:9001", "Tenant-0", "Tenant", timeout)
		// child := pidResp.Pid
		// ctx.Send(child, msg)
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
	// var rootContext = actor.EmptyRootContext
	// props := actor.PropsFromProducer(func() actor.Actor { return &Actor{} })
	// rootContext.SpawnNamed(props, "Verifier")
	console.ReadLine()
}
