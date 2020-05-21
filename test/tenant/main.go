package main

import (
	"fmt"
	"time"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"

	"Reactive-Welfare-Housing-System/src/tenant"
	// "Reactive-Welfare-Housing-System/src/shared"
)

type parentActor struct{}

func newParentActor() actor.Actor {
	return &parentActor{}
}

// func newChildActor() actor.Actor {
// 	return &tenant.Actor{}
// }

func (state *parentActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *tenant.TriggerNewApplicationRequest:
		timeout := 5 * time.Second
		pidResp, _ := remote.SpawnNamed("127.0.0.1:9001", "Tenant-0", "Tenant", timeout)
		child := pidResp.Pid
		ctx.Send(child, msg)
		// res, _ := actor.EmptyRootContext.RequestFuture(child, msg, timeout).Result()
		// response := res.(*tenant.TriggerNewApplicationResponse)
		// fmt.Printf("Response from remote %v", response.Message)

		// pidResp2, _ := remote.SpawnNamed("127.0.0.1:9001", "Tenant-1", "Tenant", timeout)
		// child2 := pidResp2.Pid
		// ctx.Send(child2, msg)
	}
}

func main() {
	// remote.Start("127.0.0.1:9000")
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
	rootContext.Send(pid, &tenant.TriggerNewApplicationRequest{})

	console.ReadLine()
}
