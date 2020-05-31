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
	case *sharedMessages.ManagerConnect:
		d.managerPID = msg.Sender
		fmt.Println("In distributor", msg.Sender)
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
				for index, occupied := range d.occupied[deleted.House.Level] {
					if occupied.HouseID == deleted.House.ID && occupied.FamilyID == deleted.FamilyID {
						d.occupied[deleted.House.Level], _ = storage.RemoveReside(d.occupied[deleted.House.Level], index)
						break
					}
				}
			} else {
				for index, vacant := range d.vacant[deleted.House.Level] {
					if vacant.HouseID == deleted.House.ID {
						d.vacant[deleted.House.Level], _ = storage.RemoveReside(d.vacant[deleted.House.Level], index)
						break
					}
				}
			}
		}
		ctx.Respond(&distributorMessages.UnqualifiedResidesACK{})
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
				d.occupied[msg.Level], _ = storage.RemoveReside(d.occupied[msg.Level], index)
				log.Print("Received HouseCheckOut ACK")
			default:
				log.Print("Received unexpected response: %#v", res)
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
				log.Print("Received HouseMatch ACK")
			case *managerMessages.HouseMatchReject:
				recv := res.(*managerMessages.HouseMatchReject)
				var occupied storage.Reside
				d.occupied[msg.Level], occupied = storage.RemoveReside(d.occupied[msg.Level], index)
				switch recv.Reason {
				case HAVEONEHOUSE:
					occupied.FamilyID = 0
					d.vacant[msg.Level] = append(d.vacant[msg.Level], occupied)
				case HOUSEMATCHED, HOUSEDONOTEXIST:
					msg.Retry = -1
					ctx.Self().Tell(&distributorMessages.MatchEmptyHouse{Request: msg})
				}
			default:
				log.Print("Received unexpected response: %#v", res)
			}
		})
	case *distributorMessages.MatchEmptyHouse:
		if len(d.vacant[msg.Request.Level]) != 0 {
			var vacant storage.Reside
			if msg.Retry != -1 {
				vacant = d.occupied[msg.Request.Level][int(msg.Retry)]
			} else {
				d.vacant[msg.Request.Level], vacant = storage.RemoveReside(d.vacant[msg.Request.Level], len(d.vacant[msg.Request.Level])-1)
				vacant.FamilyID = msg.Request.FamilyID
				d.occupied[msg.Request.Level] = append(d.occupied[msg.Request.Level], vacant)
				msg.Retry = int32(len(d.occupied[msg.Request.Level]))
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
					log.Print("Received HouseMatch ACK")
				case *managerMessages.HouseMatchReject:
					recv := res.(*managerMessages.HouseMatchReject)
					var occupied storage.Reside
					d.occupied[msg.Request.Level], occupied = storage.RemoveReside(d.occupied[msg.Request.Level], int(msg.Retry))
					switch recv.Reason {
					case HAVEONEHOUSE:
						occupied.FamilyID = 0
						d.vacant[msg.Request.Level] = append(d.vacant[msg.Request.Level], occupied)
					case HOUSEMATCHED, HOUSEDONOTEXIST:
						msg.Retry = -1
						ctx.Self().Tell(msg)
					}
				default:
					log.Print("Received unexpected response: %#v", res)
				}
			})
			return
		}
		ctx.Request(d.verifierPID, &distributorMessages.HouseApplicationReject{Request: msg.Request, Reason: "对不起，一户家庭不能租住超过一套保障房"})
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
