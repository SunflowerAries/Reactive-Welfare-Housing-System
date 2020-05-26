package manager

import (
	"Reactive-Welfare-Housing-System/src/messages/distributorMessages"
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/propertyMessages"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/messages/verifierMessages"
	"Reactive-Welfare-Housing-System/src/property"
	"Reactive-Welfare-Housing-System/src/storage"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/router"
)

type managerActor struct {
	db             storage.HouseSystem
	distributorPID *actor.PID
	propertyPID    *actor.PID
	verifierPID    *actor.PID
}

func (m *managerActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		m.propertyPID = ctx.Spawn(router.NewRoundRobinPool(5).WithProducer(property.NewPropertyActor(m.db, ctx.Self())))
	case *sharedMessages.NewHouses:
		var houses []storage.House
		for _, house := range msg.Houses {
			houses = append(houses, storage.House{Level: house.Level, Age: house.Age, Area: house.Area})
		}
		HouseID := m.db.BatchInsertHouse(houses)
		start := 0
		for i := 0; i < len(HouseID)/2; i++ {
			for j := 0; j < HouseID[2*i]; j++ {
				houses[start+j].ID = int32(HouseID[2*i+1] + j)
			}
			start += HouseID[2*i]
		}
		newhouses := make([]*sharedMessages.NewHouse, 0, len(msg.Houses))
		for _, house := range houses {
			newhouses = append(newhouses, &sharedMessages.NewHouse{ID: house.ID, Level: house.Level})
		}
		future := ctx.RequestFuture(m.distributorPID, &managerMessages.NewHouses{Houses: &sharedMessages.NewHouses{Houses: newhouses}}, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				ctx.Self().Tell(msg)
				return
			}

			switch res.(type) {
			case *distributorMessages.NewHousesACK:
				log.Print("Received Newhouse ACK")
			default:
				log.Print("Received unexpected response: %#v", res)
			}
		})
	case *sharedMessages.DistributorConnect:
		m.distributorPID = msg.Sender
		fmt.Println("In manager", msg.Sender)
	case *sharedMessages.ExaminationList:
		future := ctx.RequestFuture(m.propertyPID, &managerMessages.ExaminationList{HouseID: msg.HouseID}, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				ctx.Self().Tell(msg)
				return
			}

			switch res.(type) {
			case *propertyMessages.ExaminationACK:
				log.Print("Received Examination ACK")
			case *propertyMessages.ExaminationRejects:
				ctx.Self().Tell(res.(*propertyMessages.ExaminationRejects))
			default:
				log.Print("Received unexpected response: %#v", res)
			}
		})
	case *distributorMessages.HouseMatch:
		ret, err := m.db.InsertMatch(storage.Reside{HouseID: msg.HouseID, FamilyID: msg.FamilyID})
		if err != nil || !ret.FamilyOwnHouse {
			fmt.Printf("Insert house failed, err:%v", err)
			ctx.Respond(&managerMessages.HouseMatchReject{Match: msg})
			return
		}
		if ret.Success {
			ctx.Respond(&managerMessages.HouseMatchACK{})
			future := ctx.RequestFuture(m.verifierPID, &managerMessages.HouseMatchApprove{Match: msg}, 2000*time.Millisecond)
			ctx.AwaitFuture(future, func(res interface{}, err error) {
				if err != nil {
					ctx.Self().Tell(msg)
					return
				}
				switch res.(type) {
				case *verifierMessages.HouseMatchApproveACK:
					log.Print("Received HouseMatchApprove ACK")
				default:
					log.Print("Received unexpected response: %#v", res)
				}
			})
		} else {
			ctx.Respond(&managerMessages.HouseMatchReject{Match: msg})
			future := ctx.RequestFuture(m.verifierPID, &managerMessages.HouseMatchReject{Match: msg, Reason: "对不起，一户家庭不能租住超过一套保障房"}, 2000*time.Millisecond)
			ctx.AwaitFuture(future, func(res interface{}, err error) {
				if err != nil {
					ctx.Self().Tell(msg)
					return
				}
				switch res.(type) {
				case *verifierMessages.HouseMatchRejectACK:
					log.Print("Received HouseMatchReject ACK")
				default:
					log.Print("Received unexpected response: %#v", res)
				}
			})
		}
	case *distributorMessages.HouseCheckOut:
		err := m.db.CheckOutHouse(storage.Reside{FamilyID: msg.FamilyID, HouseID: msg.HouseID})
		if err != nil {
			fmt.Printf("Checkout house failed, err:%v", err)
			return
		}
	case *propertyMessages.ExaminationRejects:
		fmt.Println(msg)
		var houses []*propertyMessages.ExaminationReject
		var ids []int32
		for _, house := range msg.Houses {
			houses = append(houses, house)
			ids = append(ids, house.House.ID)
		}
		err := m.db.DeleteHouse(ids)
		if err != nil {
			fmt.Printf("Delete house failed, err:%v", err)
			return
		}
		ctx.Request(m.distributorPID, &managerMessages.UnqualifiedHouses{Houses: houses})
	default:
		fmt.Println(reflect.TypeOf(msg), msg)
	}
}

func NewManagerActor(db storage.HouseSystem) actor.Producer {
	return func() actor.Actor {
		return &managerActor{
			db: db,
		}
	}
}
