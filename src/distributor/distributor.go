package distributor

import (
	"Reactive-Welfare-Housing-System/src/config"
	"Reactive-Welfare-Housing-System/src/messages/distributorMessages"
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/messages/verifierMessages"
	"Reactive-Welfare-Housing-System/src/shared"
	"Reactive-Welfare-Housing-System/src/storage"
	"Reactive-Welfare-Housing-System/src/utils"
	"fmt"
	"log"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type distributorActor struct {
	db                  storage.HouseSystem
	occupied            map[int32]storage.Reside
	vacant              [config.HouseLevel + 1][]storage.Reside
	managerPID          *actor.PID
	verifierPID         *actor.PID
	newhouseCommitIndex int
}

func (d *distributorActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		d.occupied, d.vacant = d.db.InitMatchCache()
		//for index := range d.vacant {
		//	fmt.Printf("Vacant Level: %d\n", index)
		//	for _, vacant := range d.vacant[index] {
		//		fmt.Printf("%d\t", vacant.HouseID)
		//	}
		//	fmt.Printf("\n")
		//}
	case *sharedMessages.VerifierConnect:
		d.verifierPID = msg.Sender
	case *sharedMessages.ManagerConnect:
		d.managerPID = msg.Sender
		ctx.Send(d.managerPID, &sharedMessages.DistributorConnect{Sender: ctx.Self()})
	case *managerMessages.NewHouses:
		//log-based:include index in ack
		pstart := d.newhouseCommitIndex - int(msg.CommitIndex)
		for _, vacant := range msg.Houses.Houses[pstart:] {
			d.vacant[vacant.Level] = append(d.vacant[vacant.Level], storage.Reside{HouseID: vacant.ID, Level: vacant.Level})
		}
		d.newhouseCommitIndex += len(msg.Houses.Houses) - pstart
		ctx.Respond(&distributorMessages.NewHousesACK{CommitIndex: int32(d.newhouseCommitIndex)})
	case *managerMessages.UnqualifiedHouses:
		// 一次处理一条过于低效，需要优化
		for _, deleted := range msg.Houses {
			if deleted.FamilyID != 0 {
				delete(d.occupied, deleted.FamilyID)
			} else {
				for index, vacant := range d.vacant[deleted.Level] {
					if vacant.HouseID == deleted.HouseID {
						d.vacant[deleted.Level], _ = utils.RemoveReside(d.vacant[deleted.Level], index)
						break
					}
				}
			}
		}
		ctx.Respond(&distributorMessages.UnqualifiedHousesACK{})
	case *verifierMessages.HouseCheckOut:
		ctx.Respond(&distributorMessages.HouseCheckOutACK{})
		if reside, ok := d.occupied[msg.FamilyID]; ok {
			delete(d.occupied, msg.FamilyID)
			d.vacant[reside.Level] = append(d.vacant[reside.Level], reside)
			//enter the queue
			ctx.Request(d.managerPID, &distributorMessages.HouseCheckOut{FamilyID: reside.FamilyID, HouseID: reside.HouseID})
		} else {
			fmt.Printf("Family[%d] do not have any house.\n")

		}

		for _, occupied := range d.occupied[msg.Level][msg.FamilyID] {
			future := ctx.RequestFuture(d.managerPID,, 2000*time.Millisecond)
			ctx.AwaitFuture(future, func(res interface{}, err error) {
				switch res.(type) {
				case *managerMessages.HouseCheckOutACK:
					recv := res.(*managerMessages.HouseCheckOutACK)
					for i, occupy := range d.occupied[msg.Level][msg.FamilyID] {
						if occupy.HouseID == recv.HouseID {
							d.occupied[msg.Level][msg.FamilyID], _ = storage.RemoveReside(d.occupied[msg.Level][msg.FamilyID], i)
							break
						}
					}
					log.Print("Distributor: Received HouseCheckOut ACK")
				default:
					log.Print("Distributor: Received unexpected response, ", res)
				}
			})
		}
	case *verifierMessages.HouseApplicationRequest:
		for _, occupied := range d.occupied[msg.Level][msg.FamilyID] {
			if occupied.CheckOut != true {
				if msg.Retry != true {
					ctx.Respond(&distributorMessages.HouseApplicationACK{})
				}
				future := ctx.RequestFuture(d.managerPID, &distributorMessages.HouseMatch{FamilyID: occupied.FamilyID, HouseID: occupied.HouseID}, 2000*time.Millisecond)
				ctx.AwaitFuture(future, func(res interface{}, err error) {
					if err != nil {
						msg.Retry = true
						ctx.Self().Tell(msg)
						return
					}

					switch res.(type) {
					case *managerMessages.HouseMatchACK:
						log.Print("Distributor: Received HouseMatch ACK")
					case *managerMessages.HouseMatchReject:
						recv := res.(*managerMessages.HouseMatchReject)
						var occupied storage.Reside
						var i int
						for i, occupied = range d.occupied[msg.Level][msg.FamilyID] {
							if occupied.CheckOut != true {
								d.occupied[msg.Level][msg.FamilyID], _ = storage.RemoveReside(d.occupied[msg.Level][msg.FamilyID], i)
								break
							}
						}
						switch recv.Reason {
						case shared.HAVEONEHOUSE, shared.FAMILYDONOTEXIST:
							occupied.FamilyID = 0
							d.vacant[msg.Level] = append(d.vacant[msg.Level], occupied)
						case shared.HOUSEMATCHED, shared.HOUSEDONOTEXIST:
							msg.Retry = false
							ctx.Self().Tell(&distributorMessages.MatchEmptyHouse{Request: msg})
						}
					default:
						log.Print("Distributor: Received unexpected response, ", res)
					}
				})
				return
			}
		}
		msg.Retry = false
		ctx.Self().Tell(&distributorMessages.MatchEmptyHouse{Request: msg, Sender: ctx.Sender()})
	case *distributorMessages.MatchEmptyHouse:
		var index int = -1
		var vacant storage.Reside
		var i int
		for i, vacant = range d.occupied[msg.Request.Level][msg.Request.FamilyID] {
			if vacant.CheckOut != true {
				index = i
				break
			}
		}
		if index == -1 {
			if len(d.vacant[msg.Request.Level]) != 0 {
				d.vacant[msg.Request.Level], vacant = storage.RemoveReside(d.vacant[msg.Request.Level], len(d.vacant[msg.Request.Level])-1)
				vacant.FamilyID = msg.Request.FamilyID
				d.occupied[msg.Request.Level][msg.Request.FamilyID] = append(d.occupied[msg.Request.Level][msg.Request.FamilyID], vacant)
			} else {
				ctx.Request(msg.Sender, &distributorMessages.HouseApplicationReject{Request: msg.Request, Reason: "对不起，系统中暂时没有匹配房源"})
				return
			}
		}

		if msg.Request.Retry != true {
			ctx.Request(msg.Sender, &distributorMessages.HouseApplicationACK{})
		}
		future := ctx.RequestFuture(d.managerPID, &distributorMessages.HouseMatch{FamilyID: vacant.FamilyID, HouseID: vacant.HouseID}, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				msg.Request.Retry = true
				ctx.Self().Tell(msg)
				return
			}
			switch res.(type) {
			case *managerMessages.HouseMatchACK:
				log.Print("Distributor: Received HouseMatch ACK")
			case *managerMessages.HouseMatchReject:
				recv := res.(*managerMessages.HouseMatchReject)
				var occupied storage.Reside
				for i, occupied = range d.occupied[msg.Request.Level][msg.Request.FamilyID] {
					if occupied.CheckOut != true {
						d.occupied[msg.Request.Level][msg.Request.FamilyID], occupied = storage.RemoveReside(d.occupied[msg.Request.Level][msg.Request.FamilyID], i)
						break
					}
				}
				log.Print("Distributor: Received HouseMatchReject With Reason, ", recv.Reason)
				switch recv.Reason {
				case shared.HAVEONEHOUSE, shared.FAMILYDONOTEXIST:
					occupied.FamilyID = 0
					d.vacant[msg.Request.Level] = append(d.vacant[msg.Request.Level], occupied)
				case shared.HOUSEMATCHED, shared.HOUSEDONOTEXIST:
					msg.Request.Retry = false
					ctx.Self().Tell(msg)
				}
			default:
				log.Print("Distributor: Received unexpected response, ", res)
			}
		})
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
