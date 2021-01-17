package manager

import (
	"Reactive-Welfare-Housing-System/src/config"
	"Reactive-Welfare-Housing-System/src/messages/distributorMessages"
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/propertyMessages"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/messages/verifierMessages"
	"Reactive-Welfare-Housing-System/src/property"
	"Reactive-Welfare-Housing-System/src/storage"
	"Reactive-Welfare-Housing-System/src/utils"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/mailbox"
	"log"
	"reflect"
	"sort"
	"sync"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type houseLog struct {
	log         []*sharedMessages.NewHouse
	commitIndex int
	Mu          sync.RWMutex
}

type resideLog struct {
	log         []*managerMessages.UnqualifiedReside
	commitIndex int
	Mu          sync.RWMutex
}

type managerActor struct {
	db                      storage.HouseSystem
	distributorPID          *actor.PID
	propertyPID             *actor.PID
	verifierPID             *actor.PID
	newHousesLog            houseLog
	unqualifiedResides2DLog resideLog // to distributor
	unqualifiedResides2VLog resideLog // to verifier
	houseMatchIndex         utils.Index
}

func (m *managerActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		m.propertyPID = ctx.Spawn(actor.PropsFromProducer(property.NewPropertyActor(m.db, ctx.Self())).WithMailbox(mailbox.Unbounded()))
		go func() {
			for {
				select {
				// 300ms
				case <-time.Tick(config.TIME_OUT * time.Millisecond):
					go func() {
						m.unqualifiedResides2DLog.Mu.RLock()
						houseCommitIndex := m.unqualifiedResides2DLog.commitIndex
						resides := m.unqualifiedResides2DLog.log[houseCommitIndex:]
						m.unqualifiedResides2DLog.Mu.RUnlock()
						if len(resides) > 0 {
							ctx.Request(m.distributorPID, &managerMessages.UnqualifiedResides{Resides: resides, CommitIndex: int32(houseCommitIndex)})
						}
					}()

					go func() {
						m.unqualifiedResides2VLog.Mu.RLock()
						resideCommitIndex := m.unqualifiedResides2VLog.commitIndex
						resides := m.unqualifiedResides2VLog.log[resideCommitIndex:]
						m.unqualifiedResides2VLog.Mu.RUnlock()
						if len(resides) > 0 {
							ctx.Request(m.verifierPID, &managerMessages.UnqualifiedResides{Resides: resides, CommitIndex: int32(resideCommitIndex)})
						}
					}()
				}
			}
		}()
		//checkouts := m.db.ClearHouse()
		//var resides []*managerMessages.UnqualifiedReside
		//for _, checkout := range checkouts {
		//	resides = append(resides, &managerMessages.UnqualifiedReside{houseID: checkout.houseID, FamilyID: checkout.FamilyID})
		//}
		//if len(resides) != 0 {
		//	ctx.Self().Tell(&managerMessages.UnqualifiedResides{Resides: resides})
		//}
	case *sharedMessages.NewHouses:
		m.newHousesLog.Mu.Lock()
		pstart := len(m.newHousesLog.log)
		m.newHousesLog.log = append(m.newHousesLog.log, msg.Houses...)
		var houses []storage.House
		for _, house := range msg.Houses {
			houses = append(houses, storage.House{Age: house.Age, Area: house.Area, Level: house.Level})
		}
		m.newHousesLog.Mu.Unlock()
		go func(pstart int, houses []storage.House) {
			houseID := m.db.BatchInsertHouses(houses)
			idx := houseID.Idx
			pstart0 := pstart
			for i := 0; i < len(houseID.Length); i++ {
				for j := 0; j < houseID.Length[i]; j++ {
					m.newHousesLog.log[pstart+j].ID = int32(idx + j)
				}
				pstart += houseID.Length[i]
				idx += houseID.Length[i]
			}
			m.newHousesLog.Mu.RLock()
			commitIndex := m.newHousesLog.commitIndex
			m.newHousesLog.Mu.RUnlock()
			ctx.Request(m.distributorPID, &managerMessages.NewHouses{Houses: m.newHousesLog.log[commitIndex : pstart0+len(houses)], CommitIndex: int32(commitIndex)})
		}(pstart, houses)
	case *distributorMessages.NewHousesACK:
		m.newHousesLog.Mu.Lock()
		m.newHousesLog.commitIndex = utils.Max(m.newHousesLog.commitIndex, int(msg.CommitIndex))
		m.newHousesLog.Mu.Unlock()
	case *sharedMessages.DistributorConnect:
		m.distributorPID = msg.Sender
	case *sharedMessages.VerifierConnect:
		m.verifierPID = msg.Sender
	case *distributorMessages.HouseMatchRequests:
		pstart := int(msg.CommitIndex)
		var requests []storage.Reside
		m.houseMatchIndex.Mu.Lock()
		matchIndex := m.houseMatchIndex.Idx
		m.houseMatchIndex.Mu.Unlock()
		if matchIndex > pstart && matchIndex < pstart+len(msg.Requests) {
			for _, request := range msg.Requests[matchIndex-pstart:] {
				requests = append(requests, storage.Reside{HouseID: request.Match.HouseID, FamilyID: request.FamilyID, Level: request.Match.Level})
			}
			go func(requests utils.Resides) {
				sort.Sort(requests)
				unqualifieds := m.db.BatchInsertMatches(requests)
				m.houseMatchIndex.Mu.Lock()
				m.houseMatchIndex.Idx += len(requests)
				idx := m.houseMatchIndex.Idx
				m.houseMatchIndex.Mu.Unlock()

				m.unqualifiedResides2VLog.Mu.Lock()
				for _, unqualified := range unqualifieds {
					m.unqualifiedResides2VLog.log = append(m.unqualifiedResides2VLog.log, &managerMessages.UnqualifiedReside{HouseID: unqualified.HouseID, FamilyID: unqualified.FamilyID, Level: unqualified.Level})
				}
				m.unqualifiedResides2VLog.Mu.Unlock()
				ctx.Request(m.distributorPID, &managerMessages.HouseMatchRequestsACK{CommitIndex: int32(idx)})
			}(requests)
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
	case *propertyMessages.UnqualifiedHouseIDs:
		go func(houseIDs []int32) {
			rawUnqualifiedHouse := m.db.QueryReside(houseIDs)
			if err := m.db.DeleteHouseAndReside(houseIDs); err != nil {
				log.Println("propertyMessages.UnqualifiedHouseIDs[manager]: delete house and reside fail, ", err)
			}
			m.unqualifiedResides2DLog.Mu.Lock()
			m.unqualifiedResides2VLog.Mu.Lock()
			for _, house := range rawUnqualifiedHouse {
				reside := &managerMessages.UnqualifiedReside{HouseID: house.HouseID, Level: house.Level, FamilyID: house.FamilyID}
				if house.FamilyID != 0 {
					m.unqualifiedResides2VLog.log = append(m.unqualifiedResides2VLog.log, reside)
				}
				m.unqualifiedResides2DLog.log = append(m.unqualifiedResides2DLog.log, reside)
			}
			m.unqualifiedResides2VLog.Mu.Unlock()
			m.unqualifiedResides2DLog.Mu.Unlock()
		}(msg.HouseID)
	case *distributorMessages.UnqualifiedHousesACK:
		m.unqualifiedResides2DLog.Mu.Lock()
		m.unqualifiedResides2DLog.commitIndex = utils.Max(m.unqualifiedResides2DLog.commitIndex, int(msg.CommitIndex))
		m.unqualifiedResides2DLog.Mu.Unlock()
	case *verifierMessages.UnqualifiedResidesACK:
		m.unqualifiedResides2VLog.Mu.Lock()
		m.unqualifiedResides2VLog.commitIndex = utils.Max(m.unqualifiedResides2VLog.commitIndex, int(msg.CommitIndex))
		m.unqualifiedResides2VLog.Mu.Unlock()
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
