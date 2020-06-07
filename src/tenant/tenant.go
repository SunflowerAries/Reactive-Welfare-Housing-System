package tenant

import (
	"fmt"
<<<<<<< HEAD

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"

	tenantMessage "housingSystem/src/messages/tenant"
	"housingSystem/src/shared"
=======
	// console "github.com/AsynkronIT/goconsole"
	tenantMessage "Reactive-Welfare-Housing-System/src/messages/tenant"
	"Reactive-Welfare-Housing-System/src/shared"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	// _ "Reactive-Welfare-Housing-System/src/verifier"
>>>>>>> 8bd312ff8ea3b6923b202621c50548cab6706a52
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
	switch ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("Tenant Starting, initialize actor here, PID:", ctx.Self())
		t.Self = ctx.Self()
		t.verifier = actor.NewPID("127.0.0.1:8080", "Verifier")
		t.distributor = actor.NewPID("127.0.0.1:8081", "Distributor")
	case *actor.Stopped:
		fmt.Println("Tenant Stopped, actor and its children are stopped")
	case *TriggerNewApplicationRequest:
		fmt.Println("Tenant Ready to request.")
<<<<<<< HEAD
		newMsg := tenantMessage.HouseApplicationRequest{UserId: 100, UserName: "Chaokun",}
		ctx.Send(t.verifier, &newMsg)
=======
		ctx.Send(t.verifier, &tenantMessage.HouseApplicationRequest{
			UserId:   1,
			UserName: "Chaokun",
		})

>>>>>>> 8bd312ff8ea3b6923b202621c50548cab6706a52
	case *tenantMessage.HouseApplicationResponse:
		fmt.Println("The application submitted to ", ctx.Sender())
	}
}

func newTenantActor() actor.Actor {
	return &Actor{}
}

func init() {
	remote.Start("127.0.0.1:9001")
	remote.Register("Tenant", actor.PropsFromProducer(newTenantActor))
}
