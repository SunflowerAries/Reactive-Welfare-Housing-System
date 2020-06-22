package tenant

import (
	"fmt"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/messages/tenantMessages"
	"Reactive-Welfare-Housing-System/src/shared"

	"github.com/ChaokunChang/protoactor-go/actor"
	// "github.com/ChaokunChang/protoactor-go/remote"
)

// Actor The Tenant Actor
type Actor struct {
	shared.BaseActor
	verifier    *actor.PID
	distributor *actor.PID
}

// TriggerNewApplicationRequest A message that will trigger a new application
type TriggerNewApplicationRequest struct {
}

// TriggerNewApplicationResponse A message that will response trigger a new application
type TriggerNewApplicationResponse struct {
	Message string
}

// Receive Receive function for Tenant Actor
func (t *Actor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		t.Self = ctx.Self()
		t.verifier = actor.NewPID("127.0.0.1:9001", "Verifier")
	case *actor.Stopped:
		fmt.Println("Tenant Stopped, actor and its children are stopped")
	case *TriggerNewApplicationRequest:
		fmt.Println("Tenant Ready to request.")
		ctx.Send(t.verifier, &sharedMessages.NewRequest{FamilyID: 313, Level: 1})
	case *tenantMessages.HouseApplicationResponse:
		fmt.Println("The application submitted to ", ctx.Sender())
	default:
		fmt.Printf("Unexpected message for Tenant: %+v\n", msg)
	}
}

// NewTenantActor Actor Creator
func NewTenantActor() actor.Actor {
	return &Actor{}
}
