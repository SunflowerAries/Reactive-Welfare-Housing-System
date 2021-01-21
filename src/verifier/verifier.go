package verifier

import (
	"Reactive-Welfare-Housing-System/src/config"
	"Reactive-Welfare-Housing-System/src/utils"
	"fmt"
	"log"
	"sync"
	"time"

	// verifierMessage "Reactive-Welfare-Housing-System/src/messages/verifier"
	"Reactive-Welfare-Housing-System/src/messages/distributorMessages"
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/messages/verifierMessages"
	"Reactive-Welfare-Housing-System/src/shared"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type checkoutLog struct {
	log 			[]*verifierMessages.HouseCheckOut
	commitIndex		int
	Mu 				sync.RWMutex
}
// Actor The  Verifier Actor
type verifierActor struct {
	shared.BaseActor
	mamagerPID     	*actor.PID
	distributorPID 	*actor.PID
	governmentPID  	*actor.PID
	newcheckoutLog	checkoutLog
}

// Receive Receive function for Verifier Actor
func (v *verifierActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		go func() {
			select {
			//300ms
				case <- time.Tick(config.TIME_OUT * time.Millisecond):
					go func() {
						v.newcheckoutLog.Mu.RLock()
						checkoutCommitIndex := v.newcheckoutLog.commitIndex
						checkouts := v.newcheckoutLog.log[checkoutCommitIndex:]
						v.newcheckoutLog.Mu.RUnlock()
						if len(checkouts) > 0 {
							ctx.Request(v.distributorPID, &verifierMessages.HouseCheckOuts{
								Checkouts: checkouts,
								CommitIndex: int32(checkoutCommitIndex)
							})
						}
					}()
			}
		}()
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
		v.newcheckoutLog.Mu.Lock()
		pstart := len(v.newcheckoutLog.log)
		//v.newcheckoutLog.log = append(v.newcheckoutLog.log, msg.CheckOuts...)
		var checkouts []verifierMessages.HouseCheckOut
		for _, checkout := range msg.CheckOuts {
			v.newcheckoutLog.log = append(v.newcheckoutLog.log, verifierMessages.HouseCheckOut{
				FamilyID: 	checkout.FamilyID,
				Level: 	  	checkout.Level,
			})
			checkouts = append(checkouts, verifierMessages.HouseCheckOut{
				UserInfo: nil,
				FamilyID: checkout.FamilyID,
				Level:    checkout.Level,
			})
		}
		v.newcheckoutLog.Mu.Unlock()

		pstart0 := pstart
		v.newcheckoutLog.Mu.RLock()
		commitIndex := v.newcheckoutLog.commitIndex
		v.newcheckoutLog.Mu.RUnlock()
		ctx.Request(v.distributorPID, &verifierMessages.HouseCheckOuts{
			Checkouts:   v.newcheckoutLog.log[commitIndex : pstart0 + len(checkouts)],
			CommitIndex: int32(commitIndex),
		})

	case *distributorMessages.HouseCheckOutACK:
		v.newcheckoutLog.Mu.Lock()
		v.newcheckoutLog.commitIndex = utils.Max(v.newcheckoutLog.commitIndex, int(msg.CommitIndex))
		v.newcheckoutLog.Mu.Unlock()

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
		// D->V should not be high level???
		log.Print("Verifier: Received UnqualifiedResides, ", msg.Houses)
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
