package main

import (
	"fmt"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

type Hello struct{ Who string }
type ParentActor struct{}

func (state *ParentActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		props := actor.PropsFromProducer(NewChildActor)
		child := context.Spawn(props)
		context.Send(child, msg)
	}
}

func NewParentActor() actor.Actor {
	return &ParentActor{}
}

type ChildActor struct{}

func (state *ChildActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		fmt.Println("Started, initialize actor here")
	case *actor.Stopping:
		fmt.Println("Stopping, actor is about shut down")
	case *actor.Stopped:
		fmt.Println("Stopped, actor and its children are stopped")
	case *actor.Restarting:
		fmt.Println("Restarting, actor is about restart")
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
		panic("Ouch")
	}
}

func NewChildActor() actor.Actor {
	return &ChildActor{}
}

func main() {
	decider := func(reason interface{}) actor.Directive {
		fmt.Println("handling failure for child")
		return actor.StopDirective
	}
	supervisor := actor.NewOneForOneStrategy(10, 100, decider)

	context := actor.EmptyRootContext
	props := actor.
		FromProducer(NewParentActor).
		WithSupervisor(supervisor)

	pid := context.Spawn(props)
	context.Send(pid, Hello{Who: "Roger"})

	console.ReadLine()
}
