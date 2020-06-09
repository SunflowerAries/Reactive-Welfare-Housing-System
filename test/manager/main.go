package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/ChaokunChang/protoactor-go/actor"
	"github.com/ChaokunChang/protoactor-go/mailbox"
	"github.com/ChaokunChang/protoactor-go/remote"

	"Reactive-Welfare-Housing-System/src/manager"
	"Reactive-Welfare-Housing-System/src/shared"
	"Reactive-Welfare-Housing-System/src/storage"
)

type managerSupervisor struct {
	manager *actor.PID
}

func newManagerSupervisor() actor.Actor {
	return &managerSupervisor{}
}

var rootContext = actor.EmptyRootContext

func (state *managerSupervisor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("ManagerSupervisor Started, PID:", ctx.Self())
		db, err := storage.InitDB()
		if err != nil {
			log.Panic(err)
		}
		state.manager, _ = rootContext.SpawnNamed(
			actor.PropsFromProducer(manager.NewManagerActor(db)).WithMailbox(mailbox.Unbounded()),
			"Manager")
		fmt.Println("Manager Started, PID:", state.manager)
	default:
		fmt.Printf("Unexpected message for ManagerSupervisor: %+v\n", msg)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	remote.Start("127.0.0.1:9003")
	decider := func(reason interface{}) actor.Directive {
		fmt.Println("handling failure for child")
		return actor.StopDirective
	}
	supervisor := actor.NewOneForOneStrategy(10, 1000, decider)
	props := actor.
		PropsFromProducer(newManagerSupervisor).
		WithSupervisor(supervisor).
		WithMailbox(mailbox.Unbounded())

	pid, _ := rootContext.SpawnNamed(props, "ManagerSupervisor")
	shared.Use(pid)
	text, _ := reader.ReadString('\n')
	shared.Use((text))
}
