package verifier

import (
	"fmt"
	"log"
	"time"

	// verifierMessage "Reactive-Welfare-Housing-System/src/messages/verifier"
	"Reactive-Welfare-Housing-System/src/messages/distributorMessages"
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/messages/verifierMessages"
	"Reactive-Welfare-Housing-System/src/shared"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// Actor The  Verifier Actor
type verifierActor struct {
	shared.BaseActor
	mamager        *actor.PID
	distributorPID *actor.PID
	government     *actor.PID
}

// Receive Receive function for Verifier Actor
func (v *verifierActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	// case *actor.Started:
	// 	fmt.Println("Verifier Starting, initialize actor here, PID:", ctx.Self())
	// t.distributor = actor.NewPID("127.0.0.1:8081", "Distributor")
	// t.government = actor.NewPID("127.0.0.1:8082", "Government")
	// shared.Use(verifierActor)
	// case *actor.Stopped:
	// 	fmt.Println("Verifier Stopped, actor and its children are stopped")
	// case *tenantMessage.HouseApplicationRequest:
	// 	fmt.Println("Verifier Received a application from ", msg.UserName)
	// 	ctx.Send(ctx.Sender(), &tenantMessage.HouseApplicationResponse{})
	case *sharedMessages.DistributorConnect:
		v.distributorPID = msg.Sender
		ctx.Send(v.distributorPID, &sharedMessages.VerifierConnect{Sender: ctx.Self()})
	case *sharedMessages.NewRequest:
		future := ctx.RequestFuture(v.distributorPID, &verifierMessages.HouseApplicationRequest{FamilyID: msg.FamilyID, Level: msg.Level, Retry: false}, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				ctx.Self().Tell(msg)
			}

			switch res.(type) {
			case *distributorMessages.HouseApplicationACK:
				log.Print("Verifier: Received HouseApplication ACK")
			case *distributorMessages.HouseApplicationReject:
				log.Print("Verifier: Received HouseApplication Reject, ", res.(*distributorMessages.HouseApplicationReject).Reason)
			default:
				log.Print("Verifier: Request Received unexpected response, ", res)
			}
		})
	case *sharedMessages.NewRequests:
		for _, request := range msg.Requests {
			ctx.Self().Tell(request)
		}
	case *sharedMessages.NewCheckOut:
		future := ctx.RequestFuture(v.distributorPID, &verifierMessages.HouseCheckOut{FamilyID: msg.FamilyID, Level: msg.Level, Retry: false}, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				ctx.Self().Tell(msg)
			}

			switch res.(type) {
			case *distributorMessages.HouseCheckOutACK:
				log.Print("Verifier: Received HouseCheckOut ACK")
			default:
				log.Print("Verifier: CheckOut Received unexpected response, ", res)
			}
		})
	case *sharedMessages.NewCheckOuts:
		for _, checkout := range msg.CheckOuts {
			ctx.Self().Tell(checkout)
		}
	case *managerMessages.HouseMatchApprove:
		log.Print("Verifier: Received HouseMatchApprove, ", msg.Match)
		ctx.Respond(&verifierMessages.HouseMatchApproveACK{})
	case *managerMessages.HouseMatchReject:
		log.Print("Verifier: Received HouseMatchReject, ", msg.Match, msg.Reason)
		ctx.Respond(&verifierMessages.HouseMatchRejectACK{})
	case *managerMessages.HouseCheckOut:
		log.Print("Verifier: Received HouseCheckOut, ", msg.CheckOut)
		ctx.Respond(&verifierMessages.HouseCheckOutACK{})
	case *managerMessages.UnqualifiedResides:
		log.Print("Verifier: Received UnqualifiedResides, ", msg.Resides)
		ctx.Respond(&verifierMessages.UnqualifiedResidesACK{})
	default:
		fmt.Printf("%+v\n", msg)
	}
}

func NewVerifierActor() actor.Producer {
	return func() actor.Actor {
		return &verifierActor{}
	}
}

func init() {
	// remote.Start("127.0.0.1:8080")
	// remote.Register("Verifier", actor.PropsFromProducer(newVerifierActor))
	// var rootContext = actor.EmptyRootContext
	// props := actor.PropsFromProducer(func() actor.Actor { return &Actor{} })
	// rootContext.SpawnNamed(props, "Verifier")
}
