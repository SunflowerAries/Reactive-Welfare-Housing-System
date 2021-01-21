package distributor

import (
	"Reactive-Welfare-Housing-System/src/config"
	"Reactive-Welfare-Housing-System/src/messages/distributorMessages"
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/messages/verifierMessages"
	"Reactive-Welfare-Housing-System/src/storage"
	"Reactive-Welfare-Housing-System/src/utils"
	"fmt"
	"sync"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type applicationResponseLog struct {
	log         []*distributorMessages.HouseApplicationResponse
	commitIndex int
	mu          sync.RWMutex
}

type houseMatchRequestLog struct {
	log         []*distributorMessages.HouseMatchRequest
	commitIndex int
	mu          sync.RWMutex
}

type houseCheckOutLog struct {
	log         []*distributorMessages.HouseCheckOut
	commitIndex int
	mu          sync.RWMutex
}

type distributorActor struct {
	db                     storage.HouseSystem
	occupied               map[int32]storage.Reside
	vacant                 [config.HouseLevel + 1][]storage.Reside
	managerPID             *actor.PID
	verifierPID            *actor.PID
	newhouseCommitIndex    int
	applicationResponseLog applicationResponseLog
	houseMatchRequestLog   houseMatchRequestLog
	houseCheckOutLog       houseCheckOutLog
}

func (d *distributorActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		d.occupied, d.vacant = d.db.InitMatchCache()
		go func() {
			for {
				select {
				// 300ms
				case <-time.Tick(config.TIME_OUT * time.Millisecond):
					go func() {
						d.houseMatchRequestLog.mu.RLock()
						requestIndex := d.houseMatchRequestLog.commitIndex
						requests := d.houseMatchRequestLog.log[requestIndex:]
						d.houseMatchRequestLog.mu.RUnlock()
						if len(requests) > 0 {
							ctx.Request(d.managerPID, &distributorMessages.HouseMatchRequests{Requests: requests, CommitIndex: int32(requestIndex)})
						}
					}()

					go func() {
						d.applicationResponseLog.mu.RLock()
						responseIndex := d.applicationResponseLog.commitIndex
						requests0 := d.applicationResponseLog.log[responseIndex:]
						d.applicationResponseLog.mu.RUnlock()
						if len(requests0) > 0 {
							ctx.Request(d.verifierPID, &distributorMessages.HouseApplicationResponses{Responses: requests0, CommitIndex: int32(responseIndex)})
						}
					}()

					go func() {
						d.houseCheckOutLog.mu.RLock()
						checkoutIndex := d.houseCheckOutLog.commitIndex
						requests1 := d.houseCheckOutLog.log[checkoutIndex:]
						d.houseCheckOutLog.mu.RUnlock()
						if len(requests1) > 0 {
							ctx.Request(d.managerPID, &distributorMessages.HouseCheckOuts{
								Checkouts:   requests1,
								CommitIndex: int32(checkoutIndex),
							})
						}
					}()
				}
			}
		}()
	case *sharedMessages.VerifierConnect:
		d.verifierPID = msg.Sender
	case *sharedMessages.ManagerConnect:
		d.managerPID = msg.Sender
		ctx.Send(d.managerPID, &sharedMessages.DistributorConnect{Sender: ctx.Self()})
	case *managerMessages.NewHouses:
		//log-based:include index in ack
		if pstart := d.newhouseCommitIndex - int(msg.CommitIndex); pstart >= 0 && pstart < len(msg.Houses) {
			for _, vacant := range msg.Houses[pstart:] {
				d.vacant[vacant.Level] = append(d.vacant[vacant.Level], storage.Reside{HouseID: vacant.ID, Level: vacant.Level})
			}
			d.newhouseCommitIndex += len(msg.Houses) - pstart
			ctx.Respond(&distributorMessages.NewHousesACK{CommitIndex: int32(d.newhouseCommitIndex)})
		}
	case *managerMessages.UnqualifiedResides:
		// 一次处理一条过于低效，需要优化
		for _, deleted := range msg.Resides {
			if deleted.FamilyID != 0 {
				delete(d.occupied, deleted.FamilyID)
			} else {
				for index, vacant := range d.vacant[deleted.Level] {
					if vacant.HouseID == deleted.HouseID {
						d.vacant[deleted.Level], _ = storage.RemoveReside(d.vacant[deleted.Level], index)
						break
					}
				}
			}
		}
		ctx.Respond(&distributorMessages.UnqualifiedHousesACK{CommitIndex: msg.CommitIndex + int32(len(msg.Resides))})
	case *verifierMessages.HouseCheckOuts:
		for _, checkout := range msg.Checkouts {
			if reside, ok := d.occupied[checkout.FamilyID]; ok {
				d.houseCheckOutLog.log = append(d.houseCheckOutLog.log, &distributorMessages.HouseCheckOut{
					FamilyID: checkout.FamilyID,
					HouseID:  d.occupied[checkout.FamilyID].HouseID,
					Level:    d.occupied[checkout.FamilyID].Level,
				})
				delete(d.occupied, checkout.FamilyID)
				d.vacant[reside.Level] = append(d.vacant[reside.Level], reside)
			} else {
				fmt.Printf("Family[%d] do not have any house.\n")
			}
		}
		ctx.Respond(&distributorMessages.HouseCheckOutACK{
			CommitIndex: msg.CommitIndex + int32(len(msg.Checkouts)),
		})
	case *managerMessages.HouseCheckOutACK:
		d.houseCheckOutLog.mu.Lock()
		d.houseCheckOutLog.commitIndex = utils.Max(d.houseCheckOutLog.commitIndex, int(msg.CommitIndex))
		d.houseCheckOutLog.mu.Unlock()
	case *verifierMessages.HouseApplicationRequest:
		var status int32
		var match distributorMessages.HouseMatch
		if house, ok := d.occupied[msg.FamilyID]; !ok {
			if len(d.vacant[msg.Level]) > 0 {
				d.vacant[msg.Level], house = storage.RemoveReside(d.vacant[msg.Level], 0)
				house.FamilyID = msg.FamilyID
				d.occupied[msg.FamilyID] = house
				status = config.SUCCESS
				match = distributorMessages.HouseMatch{HouseID: house.HouseID, Level: house.Level}
				d.houseMatchRequestLog.log = append(d.houseMatchRequestLog.log, &distributorMessages.HouseMatchRequest{FamilyID: msg.FamilyID, Match: &match})
			} else {
				status = config.DO_NOT_HAVE_EMPTY_HOUSE
			}
		} else {
			status = config.ALREADY_HAVE_HOUSE
			match = distributorMessages.HouseMatch{HouseID: house.HouseID, Level: house.Level}
		}
		d.applicationResponseLog.log = append(d.applicationResponseLog.log, &distributorMessages.HouseApplicationResponse{FamilyID: msg.FamilyID, Status: status, Match: &match})
	case *managerMessages.HouseMatchRequestsACK:
		d.houseMatchRequestLog.mu.Lock()
		d.houseMatchRequestLog.commitIndex = utils.Max(d.houseMatchRequestLog.commitIndex, int(msg.CommitIndex))
		d.houseMatchRequestLog.mu.Unlock()
	default:
		fmt.Printf("%+v\n", msg)
	}
}

func NewDistributorActor(db storage.HouseSystem) actor.Producer {
	return func() actor.Actor {
		return &distributorActor{
			db: db,
		}
	}
}
