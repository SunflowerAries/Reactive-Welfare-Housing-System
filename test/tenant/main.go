package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ChaokunChang/protoactor-go/actor"
	"github.com/ChaokunChang/protoactor-go/mailbox"
	"github.com/ChaokunChang/protoactor-go/remote"

	"Reactive-Welfare-Housing-System/src/shared"
	"Reactive-Welfare-Housing-System/src/tenant"
	// "Reactive-Welfare-Housing-System/src/messages/sharedMessages"
)

type tenantSupervisor struct {
	tenant *actor.PID
}

func newTenantSupervisor() actor.Actor {
	return &tenantSupervisor{}
}

var rootContext = actor.EmptyRootContext

func (state *tenantSupervisor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("TenantSupervisor Started, PID:", ctx.Self())
		state.tenant, _ = rootContext.SpawnNamed(
			actor.PropsFromProducer(tenant.NewTenantActor).WithMailbox(mailbox.Unbounded()),
			"Tenant-0")
		fmt.Println("Tenant-0 Started, PID:", state.tenant)
	case *tenant.TriggerNewApplicationRequest:
		ctx.Send(state.tenant, &tenant.TriggerNewApplicationRequest{})
	default:
		fmt.Printf("Unexpected message for TenantSupervisor: %+v\n", msg)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	remote.Start("127.0.0.1:9000")

	decider := func(reason interface{}) actor.Directive {
		fmt.Println("handling failure for child")
		return actor.StopDirective
	}
	supervisor := actor.NewOneForOneStrategy(10, 1000, decider)
	props := actor.
		PropsFromProducer(newTenantSupervisor).
		WithSupervisor(supervisor).
		WithMailbox(mailbox.Unbounded())

	pid, _ := rootContext.SpawnNamed(props, "TenantSupervisor")
	rootContext.Send(pid, &tenant.TriggerNewApplicationRequest{})

	text, _ := reader.ReadString('\n')
	shared.Use((text))
}
