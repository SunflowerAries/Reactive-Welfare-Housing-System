package manager

import (
	"Reactive-Welfare-Housing-System/src/distributor"
	"Reactive-Welfare-Housing-System/src/messages/distributorMessages"
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/propertyMessages"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/messages/verifierMessages"
	"Reactive-Welfare-Housing-System/src/property"
	"Reactive-Welfare-Housing-System/src/storage"
	"database/sql"
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
		fmt.Print(msg)
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
			log.Print("Insert house failed, err:%v", err)
			var reason int32
			if ret.HouseMatched {
				reason = distributor.HOUSEMATCHED
			} else if err == sql.ErrNoRows {
				reason = distributor.HOUSEDONOTEXIST
			} else {
				log.Panic(err)
			}
			ctx.Respond(&managerMessages.HouseMatchReject{Match: msg, Reason: reason})
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
			if ret.FamilyOwnHouse != true {
				log.Panic(ret)
			}
			ctx.Respond(&managerMessages.HouseMatchReject{Match: msg, Reason: distributor.HAVEONEHOUSE})
			future := ctx.RequestFuture(m.verifierPID, &managerMessages.HouseMatchReject{Match: msg, Reason: distributor.HAVEONEHOUSE}, 2000*time.Millisecond)
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
		if msg.Retry != true {
			err := m.db.CheckOutHouse(storage.Reside{FamilyID: msg.FamilyID, HouseID: msg.HouseID})
			if err != nil {
				log.Print("Checkout house failed, err:%v", err)
				return
			}
			ctx.Respond(&managerMessages.HouseCheckOutACK{})
		}
		future := ctx.RequestFuture(m.verifierPID, msg, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				ctx.Self().Tell(msg)
				return
			}

			switch res.(type) {
			case *verifierMessages.HouseCheckOutACK:
				m.db.DeleteReside(storage.Reside{FamilyID: msg.FamilyID, HouseID: msg.HouseID})
				log.Print("Received UnqualifiedResides ACK")
			default:
				log.Print("Received unexpected response: %#v", res)
			}
		})
	case *propertyMessages.ExaminationRejects:
		var updatedhouses, houses []*propertyMessages.ExaminationReject
		var updatedids, deletedids []int32
		for _, house := range msg.Houses {
			if house.FamilyID != 0 {
				updatedhouses = append(updatedhouses, house)
				updatedids = append(updatedids, house.House.ID)
			} else {
				deletedids = append(deletedids, house.House.ID)
			}
			houses = append(houses, house)
		}
		if len(updatedids) != 0 {
			err := m.db.SignedDeleteHouse(updatedids)
			if err != nil {
				log.Print("Update house failed, err:%v", err)
			}
		}
		if len(deletedids) != 0 {
			err := m.db.DeleteHouse(deletedids)
			if err != nil {
				log.Print("Delete house failed, err:%v", err)
			}
		}

		ctx.Self().Tell(&managerMessages.UnqualifiedResides{Resides: updatedhouses})
		ctx.Self().Tell(&managerMessages.UnqualifiedHouses{Houses: houses})
	case *managerMessages.UnqualifiedHouses:
		fmt.Println(msg)
		future := ctx.RequestFuture(m.distributorPID, msg, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				ctx.Self().Tell(msg)
				return
			}

			switch res.(type) {
			case *distributorMessages.UnqualifiedResidesACK:
				log.Print("Received UnqualifiedResides ACK")
			default:
				log.Print("Received unexpected response: %#v", res)
			}
		})
	case *managerMessages.UnqualifiedResides:
		fmt.Println(msg)
		future := ctx.RequestFuture(m.verifierPID, msg, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				ctx.Self().Tell(msg)
				return
			}

			switch res.(type) {
			case *verifierMessages.UnqualifiedResidesACK:
				log.Print("Received UnqualifiedResides ACK")
				var ids []int32
				for _, house := range msg.Resides {
					ids = append(ids, house.House.ID)
				}
				err := m.db.DeleteHouse(ids)
				if err != nil {
					log.Print("Delete house failed, err:%v", err)
				}
			default:
				log.Print("Received unexpected response: %#v", res)
			}
		})
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
