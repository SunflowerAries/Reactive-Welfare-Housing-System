package main

import (
	"bufio"
	"fmt"
	"os"
	"log"

	"github.com/ChaokunChang/protoactor-go/actor"
	"github.com/ChaokunChang/protoactor-go/mailbox"
	"github.com/ChaokunChang/protoactor-go/remote"

	"Reactive-Welfare-Housing-System/src/distributor"
	"Reactive-Welfare-Housing-System/src/shared"
	"Reactive-Welfare-Housing-System/src/storage"
)

type distributorSupervisor struct {
	distributor *actor.PID
}

func newDistributorSupervisor() actor.Actor {
	return &distributorSupervisor{}
}

var rootContext = actor.EmptyRootContext

func (state *distributorSupervisor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("DistributorSupervisor Started, PID:", ctx.Self())
		db, err := storage.InitDB()
		if err != nil {
			log.Panic(err)
		}
		state.distributor, _ = rootContext.SpawnNamed(
			actor.PropsFromProducer(distributor.NewDistributorActor(db)).WithMailbox(mailbox.Unbounded()),
			"Distributor")
		fmt.Println("Distributor Started, PID:", state.distributor)
	default:
		fmt.Printf("Unexpected message for DistributorSupervisor: %+v\n", msg)
	}
}

func main() {
	remote.Start("127.0.0.1:9002")
	reader := bufio.NewReader(os.Stdin)
	decider := func(reason interface{}) actor.Directive {
		fmt.Println("handling failure for child")
		return actor.StopDirective
	}
	supervisor := actor.NewOneForOneStrategy(10, 1000, decider)
	props := actor.
		PropsFromProducer(newDistributorSupervisor).
		WithSupervisor(supervisor).
		WithMailbox(mailbox.Unbounded())

	pid, _ := rootContext.SpawnNamed(props, "DistributorSupervisor")
	shared.Use(pid)
	text, _ := reader.ReadString('\n')
	shared.Use((text))
}
