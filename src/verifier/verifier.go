package verifier

import (
	"fmt"
	tenantMessage "housingSystem/src/messages/tenant"
	// verifierMessage "housingSystem/src/messages/verifier"
	"housingSystem/src/shared"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
)

// Actor The  Verifier Actor
type Actor struct {
	shared.BaseActor
	mamager     *actor.PID
	distributor *actor.PID
	government  *actor.PID
}

// Receive Receive function for Verifier Actor
func (t *Actor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("Verifier Starting, initialize actor here, PID:", ctx.Self())
		t.Self = ctx.Self()
		// t.distributor = actor.NewPID("127.0.0.1:8081", "Distributor")
		// t.government = actor.NewPID("127.0.0.1:8082", "Government")
		// shared.Use(verifierActor)
	case *actor.Stopped:
		fmt.Println("Verifier Stopped, actor and its children are stopped")
	case *tenantMessage.HouseApplicationRequest:
		fmt.Println("Verifier Received a application from ", msg.UserName)
		ctx.Send(ctx.Sender(), &tenantMessage.HouseApplicationResponse{})
	}
}

func newVerifierActor() actor.Actor {
	return &Actor{}
}

func init() {
	remote.Start("127.0.0.1:8080")
	remote.Register("Verifier", actor.PropsFromProducer(newVerifierActor))
	// var rootContext = actor.EmptyRootContext
	// props := actor.PropsFromProducer(func() actor.Actor { return &Actor{} })
	// rootContext.SpawnNamed(props, "Verifier")
}
