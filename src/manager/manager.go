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

type houseLog struct {
	log         []storage.House
	commitIndex index
}

type resideLog struct {
	log         []storage.Reside
	commitIndex index
}

type managerActor struct {
	db                    storage.HouseSystem
	distributorPID        *actor.PID
	propertyPID           *actor.PID
	verifierPID           *actor.PID
	newHousesLog          houseLog
	unqualifiedHousesLog  houseLog
	unqualifiedResidesLog resideLog
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
		pstart := len(m.newHousesLog.log)
		for _, house := range msg.Houses {
			m.newHousesLog.log = append(m.newHousesLog.log, storage.House{Level: house.Level, Age: house.Age, Area: house.Area})
		}
		go func(pstart int, length int) {
			houseID := m.db.BatchInsertHouse(m.newHousesLog.log[pstart : pstart+length])
			for i := 0; i < len(houseID); i++ {
				for j := 0; j < houseID[i].Length; j++ {
					m.newHousesLog.log[pstart+j].ID = int32(houseID[i].Idx + j)
				}
				pstart += houseID[i].Length
			}
			m.newHousesLog.commitIndex.mu.Lock()
			commitIndex := m.newHousesLog.commitIndex.idx
			m.newHousesLog.commitIndex.mu.Unlock()
			newhouses := make([]*sharedMessages.NewHouse, 0, len(m.newHousesLog.log)-commitIndex)
			for _, house := range m.newHousesLog.log[commitIndex:] {
				newhouses = append(newhouses, &sharedMessages.NewHouse{ID: house.ID, Level: house.Level})
			}
			ctx.Request(m.distributorPID, &managerMessages.NewHouses{Houses: &sharedMessages.NewHouses{Houses: newhouses}, CommitIndex: int32(commitIndex)})
		}(pstart, len(m.newHousesLog.log)-pstart)
	case *distributorMessages.NewHousesACK:
		m.newHousesLog.commitIndex.mu.Lock()
		m.newHousesLog.commitIndex.idx = utils.Max(m.newHousesLog.commitIndex.idx, int(msg.CommitIndex))
		m.newHousesLog.commitIndex.mu.Unlock()
	case *sharedMessages.DistributorConnect:
		m.distributorPID = msg.Sender
	case *sharedMessages.VerifierConnect:
		m.verifierPID = msg.Sender
	case *distributorMessages.HouseMatch:
		ret, err := m.db.InsertMatch(storage.Reside{HouseID: msg.HouseID, FamilyID: msg.FamilyID})
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
		go func(houseIDs []int32) {
			var updatedResides, houses []*managerMessages.ExaminationReject
			examinedHouse := m.db.QueryReside(houseIDs)
			if err := m.db.DeleteHouseAndReside(houseIDs); err != nil {
				log.Printf("propertyMessages.ExaminationRejects[manager]: delete house and reside fail, ", err)
			}
			for _, house := range examinedHouse {
				if house.FamilyID != 0 {
					m.unqualifiedResidesLog.log = append(m.unqualifiedResidesLog.log, storage.Reside{HouseID: house.HouseID, Level: house.Level, FamilyID: house.FamilyID})
				} else {
					m.unqualifiedHousesLog.log = append(m.unqualifiedHousesLog.log, storage.House{ID: house.HouseID, Level: house.Level})
				}
			}
			m.unqualifiedHousesLog.commitIndex.mu.Lock()
			houseCommitIndex := m.unqualifiedHousesLog.commitIndex.idx
			m.unqualifiedHousesLog.commitIndex.mu.Unlock()
			for _, house := range m.unqualifiedHousesLog.log[houseCommitIndex:] {
				houses = append(houses, &managerMessages.ExaminationReject{HouseID: house.ID, Level: house.Level})
			}
			ctx.Request(m.distributorPID, &managerMessages.UnqualifiedHouses{Houses: houses})

			m.unqualifiedResidesLog.commitIndex.mu.Lock()
			resideCommitIndex := m.unqualifiedResidesLog.commitIndex.idx
			m.unqualifiedResidesLog.commitIndex.mu.Unlock()
			for _, reside := range m.unqualifiedResidesLog.log[resideCommitIndex:] {
				updatedResides = append(updatedResides, &managerMessages.ExaminationReject{HouseID: reside.HouseID, Level: reside.Level, FamilyID: reside.FamilyID})
			}
			ctx.Request(m.verifierPID, &managerMessages.UnqualifiedResides{Resides: updatedResides})
		}(msg.HouseID)
	case *distributorMessages.UnqualifiedHousesACK:
		m.unqualifiedHousesLog.commitIndex.mu.Lock()
		m.unqualifiedHousesLog.commitIndex.idx = utils.Max(m.unqualifiedHousesLog.commitIndex.idx, int(msg.CommitIndex))
		m.unqualifiedHousesLog.commitIndex.mu.Unlock()
	case *verifierMessages.UnqualifiedResidesACK:
		m.unqualifiedResidesLog.commitIndex.mu.Lock()
		m.unqualifiedResidesLog.commitIndex.idx = utils.Max(m.unqualifiedResidesLog.commitIndex.idx, int(msg.CommitIndex))
		m.unqualifiedResidesLog.commitIndex.mu.Unlock()
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
