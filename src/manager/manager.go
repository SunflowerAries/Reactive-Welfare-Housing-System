package manager

import (
	"Reactive-Welfare-Housing-System/src/messages/distributorMessages"
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/propertyMessages"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/messages/verifierMessages"
	"Reactive-Welfare-Housing-System/src/property"
	"Reactive-Welfare-Housing-System/src/shared"
	"Reactive-Welfare-Housing-System/src/storage"
	"Reactive-Welfare-Housing-System/src/utils"
	"database/sql"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/mailbox"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type index struct {
	idx int
	mu  sync.Mutex
}

type newhouseLog struct {
	log         []storage.House
	commitIndex index
}

type managerActor struct {
	db             storage.HouseSystem
	distributorPID *actor.PID
	propertyPID    *actor.PID
	verifierPID    *actor.PID
	newhousesLog   newhouseLog
}

func (m *managerActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		m.propertyPID = ctx.Spawn(actor.PropsFromProducer(property.NewPropertyActor(m.db, ctx.Self())).WithMailbox(mailbox.Unbounded()))
		//checkouts := m.db.ClearHouse()
		//var resides []*managerMessages.ExaminationReject
		//for _, checkout := range checkouts {
		//	resides = append(resides, &managerMessages.ExaminationReject{houseID: checkout.houseID, FamilyID: checkout.FamilyID})
		//}
		//if len(resides) != 0 {
		//	ctx.Self().Tell(&managerMessages.UnqualifiedResides{Resides: resides})
		//}
	case *sharedMessages.NewHouses:
		pstart := len(m.newhousesLog.log)
		for _, house := range msg.Houses {
			m.newhousesLog.log = append(m.newhousesLog.log, storage.House{Level: house.Level, Age: house.Age, Area: house.Area})
		}
		go func(pstart int, length int) {
			houseID := m.db.BatchInsertHouse(m.newhousesLog.log[pstart : pstart+length])
			for i := 0; i < len(houseID); i++ {
				for j := 0; j < houseID[i].Length; j++ {
					m.newhousesLog.log[pstart+j].ID = int32(houseID[i].Idx + j)
				}
				pstart += houseID[i].Length
			}
			m.newhousesLog.commitIndex.mu.Lock()
			commitIndex := m.newhousesLog.commitIndex.idx
			m.newhousesLog.commitIndex.mu.Unlock()
			newhouses := make([]*sharedMessages.NewHouse, 0, len(m.newhousesLog.log)-commitIndex)
			for _, house := range m.newhousesLog.log[commitIndex:] {
				newhouses = append(newhouses, &sharedMessages.NewHouse{ID: house.ID, Level: house.Level})
			}
			ctx.Request(m.distributorPID, &managerMessages.NewHouses{Houses: &sharedMessages.NewHouses{Houses: newhouses}, CommitIndex: int32(commitIndex)})
		}(pstart, len(m.newhousesLog.log)-pstart)
	case *distributorMessages.NewHousesACK:
		m.newhousesLog.commitIndex.mu.Lock()
		m.newhousesLog.commitIndex.idx = utils.Max(m.newhousesLog.commitIndex.idx, int(msg.CommitIndex))
		m.newhousesLog.commitIndex.mu.Unlock()
	case *sharedMessages.DistributorConnect:
		m.distributorPID = msg.Sender
	case *sharedMessages.VerifierConnect:
		m.verifierPID = msg.Sender
	case *sharedMessages.ExaminationList:
		future := ctx.RequestFuture(m.propertyPID, &managerMessages.ExaminationList{houseID: msg.houseID}, 2000*time.Millisecond)
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
		ret, err := m.db.InsertMatch(storage.Reside{houseID: msg.houseID, FamilyID: msg.FamilyID})
		var reason int32
		if err != nil || (!ret.FamilyOwnHouse && !ret.Success) {
			log.Print("Manager: Insert house failed, ", err)
			if ret.HouseMatched {
				reason = shared.HOUSEMATCHED
			} else if err == sql.ErrNoRows {
				reason = shared.HOUSEDONOTEXIST
			} else {
				reason = shared.FAMILYDONOTEXIST
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
			err := m.db.CheckOutHouse(storage.Reside{FamilyID: msg.FamilyID, houseID: msg.houseID})
			if err != nil {
				log.Print("Manager: Checkout house failed ", err)
				return
			}
			ctx.Respond(&managerMessages.HouseCheckOutACK{houseID: msg.houseID})
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
				m.db.DeleteReside(storage.Reside{FamilyID: msg.FamilyID, houseID: msg.houseID})
				log.Print("Manager: Received HouseCheckOut ACK")
			default:
				log.Print("Manager: Received unexpected response, ", res)
			}
		})
	case *propertyMessages.ExaminationRejects:
		var updatedhouses, houses []*managerMessages.ExaminationReject
		ExaminedHouse := m.db.QueryReside(msg.houseID)
		var updatedids, deletedids []int32
		for _, house := range ExaminedHouse {
			if house.FamilyID != 0 {
				updatedhouses = append(updatedhouses, &managerMessages.ExaminationReject{houseID: house.houseID, Age: house.Age, Area: house.Area, Level: house.Level, FamilyID: house.FamilyID})
				updatedids = append(updatedids, house.houseID)
			} else {
				deletedids = append(deletedids, house.houseID)
			}
			houses = append(houses, &managerMessages.ExaminationReject{houseID: house.houseID, Age: house.Age, Area: house.Area, Level: house.Level, FamilyID: house.FamilyID})
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
					ids = append(ids, house.houseID)
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
