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
		checkouts := m.db.ClearHouse()
		var resides []*managerMessages.ExaminationReject
		for _, checkout := range checkouts {
			resides = append(resides, &managerMessages.ExaminationReject{HouseID: checkout.HouseID, FamilyID: checkout.FamilyID})
		}
		if len(resides) != 0 {
			ctx.Self().Tell(&managerMessages.UnqualifiedResides{Resides: resides})
		}
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
				log.Print("Manager: Received Newhouse ACK")
			default:
				log.Print("Manager: Received unexpected response, ", res)
			}
		})
	case *sharedMessages.DistributorConnect:
		m.distributorPID = msg.Sender
	case *sharedMessages.VerifierConnect:
		m.verifierPID = msg.Sender
	case *sharedMessages.ExaminationList:
		future := ctx.RequestFuture(m.propertyPID, &managerMessages.ExaminationList{HouseID: msg.HouseID}, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				ctx.Self().Tell(msg)
				return
			}

			switch res.(type) {
			case *propertyMessages.ExaminationACK:
				log.Print("Manager: Received Examination ACK")
			case *propertyMessages.ExaminationRejects:
				ctx.Self().Tell(res.(*propertyMessages.ExaminationRejects))
			default:
				log.Print("Manager: Received unexpected response, ", res)
			}
		})
	case *distributorMessages.HouseMatch:
		ret, err := m.db.InsertMatch(storage.Reside{HouseID: msg.HouseID, FamilyID: msg.FamilyID})
		var reason int32
		if err != nil || (!ret.FamilyOwnHouse && !ret.Success) {
			log.Print("Manager: Insert house failed, ", err)
			if ret.HouseMatched {
				reason = distributor.HOUSEMATCHED
			} else if err == sql.ErrNoRows {
				reason = distributor.HOUSEDONOTEXIST
			} else {
				reason = distributor.FAMILYDONOTEXIST
			}
		}
		if ret.Success && err == nil {
			ctx.Respond(&managerMessages.HouseMatchACK{})
			future := ctx.RequestFuture(m.verifierPID, &managerMessages.HouseMatchApprove{Match: msg}, 2000*time.Millisecond)
			ctx.AwaitFuture(future, func(res interface{}, err error) {
				if err != nil {
					ctx.Self().Tell(msg)
					return
				}
				switch res.(type) {
				case *verifierMessages.HouseMatchApproveACK:
					log.Print("Manager: Received HouseMatchApprove ACK")
				default:
					log.Print("Manager: Received unexpected response, ", res)
				}
			})
		} else {
			ctx.Respond(&managerMessages.HouseMatchReject{Match: msg, Reason: reason})
			future := ctx.RequestFuture(m.verifierPID, &managerMessages.HouseMatchReject{Match: msg, Reason: reason}, 2000*time.Millisecond)
			ctx.AwaitFuture(future, func(res interface{}, err error) {
				if err != nil {
					ctx.Self().Tell(msg)
					return
				}
				switch res.(type) {
				case *verifierMessages.HouseMatchRejectACK:
					log.Print("Manager: Received HouseMatchReject ACK")
				default:
					log.Print("Manager: Received unexpected response, ", res)
				}
			})
		}
	case *distributorMessages.HouseCheckOut:
		if msg.Retry != true {
			err := m.db.CheckOutHouse(storage.Reside{FamilyID: msg.FamilyID, HouseID: msg.HouseID})
			if err != nil {
				log.Print("Manager: Checkout house failed ", err)
				return
			}
			ctx.Respond(&managerMessages.HouseCheckOutACK{HouseID: msg.HouseID})
		}
		future := ctx.RequestFuture(m.verifierPID, &managerMessages.HouseCheckOut{CheckOut: msg}, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				msg.Retry = true
				ctx.Self().Tell(msg)
				return
			}

			switch res.(type) {
			case *verifierMessages.HouseCheckOutACK:
				m.db.DeleteReside(storage.Reside{FamilyID: msg.FamilyID, HouseID: msg.HouseID})
				log.Print("Manager: Received HouseCheckOut ACK")
			default:
				log.Print("Manager: Received unexpected response, ", res)
			}
		})
	case *propertyMessages.ExaminationRejects:
		var updatedhouses, houses []*managerMessages.ExaminationReject
		ExaminedHouse := m.db.QueryReside(msg.HouseID)
		var updatedids, deletedids []int32
		for _, house := range ExaminedHouse {
			if house.FamilyID != 0 {
				updatedhouses = append(updatedhouses, &managerMessages.ExaminationReject{HouseID: house.HouseID, Age: house.Age, Area: house.Area, Level: house.Level, FamilyID: house.FamilyID})
				updatedids = append(updatedids, house.HouseID)
			} else {
				deletedids = append(deletedids, house.HouseID)
			}
			houses = append(houses, &managerMessages.ExaminationReject{HouseID: house.HouseID, Age: house.Age, Area: house.Area, Level: house.Level, FamilyID: house.FamilyID})
		}
		if len(updatedids) != 0 {
			err := m.db.SignedDeleteHouse(updatedids)
			if err != nil {
				log.Print("Manager: Update house failed, ", err)
			}
		}
		if len(deletedids) != 0 {
			err := m.db.DeleteHouse(deletedids)
			if err != nil {
				log.Print("Manager: Delete house failed, ", err)
			}
		}
		if len(houses) != 0 {
			ctx.Self().Tell(&managerMessages.UnqualifiedHouses{Houses: houses})
			if len(updatedhouses) != 0 {
				ctx.Self().Tell(&managerMessages.UnqualifiedResides{Resides: updatedhouses})
			}
		}
	case *managerMessages.UnqualifiedHouses:
		future := ctx.RequestFuture(m.distributorPID, msg, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				ctx.Self().Tell(msg)
				return
			}

			switch res.(type) {
			case *distributorMessages.UnqualifiedHousesACK:
				log.Print("Manager: Received UnqualifiedHouses ACK")
			default:
				log.Print("Manager: Received unexpected response, ", res)
			}
		})
	case *managerMessages.UnqualifiedResides:
		future := ctx.RequestFuture(m.verifierPID, msg, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				ctx.Self().Tell(msg)
				return
			}

			switch res.(type) {
			case *verifierMessages.UnqualifiedResidesACK:
				log.Print("Manager: Received UnqualifiedResides ACK")
				var ids []int32
				for _, house := range msg.Resides {
					ids = append(ids, house.HouseID)
				}
				err := m.db.DeleteHouse(ids)
				if err != nil {
					log.Print("Manager: Delete house failed, ", err)
				}
			default:
				log.Print("Manager: Received unexpected response, ", res)
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
