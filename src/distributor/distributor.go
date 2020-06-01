package distributor

import (
	"Reactive-Welfare-Housing-System/src/messages/distributorMessages"
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/messages/verifierMessages"
	"Reactive-Welfare-Housing-System/src/storage"
	"fmt"
	"log"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type distributorActor struct {
	db          storage.HouseSystem
	occupied    [storage.HouseLevel + 1][]storage.Reside
	vacant      [storage.HouseLevel + 1][]storage.Reside
	managerPID  *actor.PID
	verifierPID *actor.PID
}

const HAVEONEHOUSE = 1
const HOUSEMATCHED = 2
const HOUSEDONOTEXIST = 3
const FAMILYDONOTEXIST = 4

func (d *distributorActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		d.occupied, d.vacant = d.db.InitMatchCache()
		for index := range d.occupied {
			fmt.Printf("Occupied Level: %d\n", index)
			fmt.Println("HouseID\tFamilyID")
			for _, occupied := range d.occupied[index] {
				fmt.Printf("%d\t%d\n", occupied.HouseID, occupied.FamilyID)
			}
		}
		for index := range d.vacant {
			fmt.Printf("Vacant Level: %d\n", index)
			for _, vacant := range d.vacant[index] {
				fmt.Printf("%d\t", vacant.HouseID)
			}
			fmt.Printf("\n")
		}
	case *sharedMessages.VerifierConnect:
		d.verifierPID = msg.Sender
	case *sharedMessages.ManagerConnect:
		d.managerPID = msg.Sender
		ctx.Send(d.managerPID, &sharedMessages.DistributorConnect{Sender: ctx.Self()})
	case *managerMessages.NewHouses:
		for _, vacant := range msg.Houses.Houses {
			d.vacant[vacant.Level] = append(d.vacant[vacant.Level], storage.Reside{HouseID: vacant.ID, Level: vacant.Level})
		}
		for i := 1; i < 4; i++ {
			fmt.Printf("Vacant Level: %d\n", i)
			for _, vacant := range d.vacant[i] {
				fmt.Printf("%d\t", vacant.HouseID)
			}
			fmt.Printf("\n")
		}
		ctx.Respond(&distributorMessages.NewHousesACK{})
	case *managerMessages.UnqualifiedHouses:
		// 一次处理一条过于低效，需要优化
		for _, deleted := range msg.Houses {
			if deleted.FamilyID != 0 {
				for index, occupied := range d.occupied[deleted.Level] {
					if occupied.HouseID == deleted.HouseID && occupied.FamilyID == deleted.FamilyID {
						d.occupied[deleted.Level], _ = storage.RemoveReside(d.occupied[deleted.Level], index)
						break
					}
				}
			} else {
				for index, vacant := range d.vacant[deleted.Level] {
					if vacant.HouseID == deleted.HouseID {
						d.vacant[deleted.Level], _ = storage.RemoveReside(d.vacant[deleted.Level], index)
						break
					}
				}
			}
		}
		ctx.Respond(&distributorMessages.UnqualifiedHousesACK{})
	case *verifierMessages.HouseCheckOut:
		var index = -1
		if msg.Retry != -1 {
			index = int(msg.Retry)
		} else {
			for i, occupied := range d.occupied[msg.Level] {
				if occupied.FamilyID == msg.FamilyID {
					if occupied.CheckOut != true {
						occupied.CheckOut = true
						d.vacant[msg.Level] = append(d.vacant[msg.Level], storage.Reside{HouseID: occupied.HouseID, Level: msg.Level})
					}
					index = i
					break
				}
			}
			// Choice: No Reside or ACK
			ctx.Respond(&distributorMessages.HouseCheckOutACK{})
		}

		if index == -1 {
			return
		}
		future := ctx.RequestFuture(d.managerPID, &distributorMessages.HouseCheckOut{FamilyID: d.occupied[msg.Level][index].FamilyID, HouseID: d.occupied[msg.Level][index].HouseID}, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				msg.Retry = int32(index)
				ctx.Self().Tell(msg)
				return
			}

			switch res.(type) {
			case *managerMessages.HouseCheckOutACK:
				for i, occupied := range d.occupied[msg.Level] {
					if occupied.FamilyID == msg.FamilyID && occupied.CheckOut != true {
						d.occupied[msg.Level], _ = storage.RemoveReside(d.occupied[msg.Level], i)
						break
					}
				}
				log.Print("Distributor: Received HouseCheckOut ACK")
			default:
				log.Print("Distributor: Received unexpected response, ", res)
			}
		})
	case *verifierMessages.HouseApplicationRequest:
		var index = -1
		if msg.Retry != -1 {
			index = int(msg.Retry)
		} else {
			for i, occupied := range d.occupied[msg.Level] {
				if occupied.FamilyID == msg.FamilyID && occupied.CheckOut != true {
					index = i
					break
				}
			}
			ctx.Respond(&distributorMessages.HouseApplicationACK{})
			msg.Retry = int32(index)
		}
		if index == -1 {
			ctx.Self().Tell(&distributorMessages.MatchEmptyHouse{Request: msg})
			return
		}
		future := ctx.RequestFuture(d.managerPID, &distributorMessages.HouseMatch{FamilyID: d.occupied[msg.Level][index].FamilyID, HouseID: d.occupied[msg.Level][index].HouseID}, 2000*time.Millisecond)
		ctx.AwaitFuture(future, func(res interface{}, err error) {
			if err != nil {
				ctx.Self().Tell(msg)
				return
			}

			switch res.(type) {
			case *managerMessages.HouseMatchACK:
				log.Print("Distributor: Received HouseMatch ACK")
			case *managerMessages.HouseMatchReject:
				recv := res.(*managerMessages.HouseMatchReject)
				var occupied storage.Reside
				for i, occupied := range d.occupied[msg.Level] {
					if occupied.FamilyID == msg.FamilyID && occupied.CheckOut != true {
						d.occupied[msg.Level], occupied = storage.RemoveReside(d.occupied[msg.Level], i)
						break
					}
				}
				switch recv.Reason {
				case HAVEONEHOUSE, FAMILYDONOTEXIST:
					occupied.FamilyID = 0
					d.vacant[msg.Level] = append(d.vacant[msg.Level], occupied)
				case HOUSEMATCHED, HOUSEDONOTEXIST:
					msg.Retry = -1
					ctx.Self().Tell(&distributorMessages.MatchEmptyHouse{Request: msg})
				}
			default:
				log.Print("Distributor: Received unexpected response, ", res)
			}
		})
	case *distributorMessages.MatchEmptyHouse:
		if len(d.vacant[msg.Request.Level]) != 0 {
			var vacant storage.Reside
			if msg.Request.Retry != -1 {
				vacant = d.occupied[msg.Request.Level][int(msg.Request.Retry)]
			} else {
				d.vacant[msg.Request.Level], vacant = storage.RemoveReside(d.vacant[msg.Request.Level], len(d.vacant[msg.Request.Level])-1)
				vacant.FamilyID = msg.Request.FamilyID
				d.occupied[msg.Request.Level] = append(d.occupied[msg.Request.Level], vacant)
				msg.Request.Retry = int32(len(d.occupied[msg.Request.Level]) - 1)
			}
			// 发给 manager
			future := ctx.RequestFuture(d.managerPID, &distributorMessages.HouseMatch{FamilyID: vacant.FamilyID, HouseID: vacant.HouseID}, 2000*time.Millisecond)
			ctx.AwaitFuture(future, func(res interface{}, err error) {
				if err != nil {
					ctx.Self().Tell(msg)
					return
				}

				switch res.(type) {
				case *managerMessages.HouseMatchACK:
					log.Print("Distributor: Received HouseMatch ACK")
				case *managerMessages.HouseMatchReject:
					recv := res.(*managerMessages.HouseMatchReject)
					var occupied storage.Reside
					for i, occupied := range d.occupied[msg.Request.Level] {
						if occupied.FamilyID == msg.Request.FamilyID && occupied.CheckOut != true {
							d.occupied[msg.Request.Level], occupied = storage.RemoveReside(d.occupied[msg.Request.Level], i)
							break
						}
					}
					switch recv.Reason {
					case HAVEONEHOUSE, FAMILYDONOTEXIST:
						occupied.FamilyID = 0
						d.vacant[msg.Request.Level] = append(d.vacant[msg.Request.Level], occupied)
					case HOUSEMATCHED, HOUSEDONOTEXIST:
						msg.Request.Retry = -1
						ctx.Self().Tell(msg)
					}
				default:
					log.Print("Distributor: Received unexpected response, ", res)
				}
			})
			return
		}
		ctx.Request(d.verifierPID, &distributorMessages.HouseApplicationReject{Request: msg.Request, Reason: "对不起，系统中暂时没有匹配房源"})
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
