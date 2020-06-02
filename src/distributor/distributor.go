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
	occupied    [storage.HouseLevel + 1]map[int32][]storage.Reside
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
		for i := 1; i < 4; i++ {
			d.occupied[i] = make(map[int32][]storage.Reside)
		}
		var occupieds [storage.HouseLevel + 1][]storage.Reside

		occupieds, d.vacant = d.db.InitMatchCache()

		for index := range occupieds {
			fmt.Printf("Occupied Level: %d\n", index)
			fmt.Println("HouseID\tFamilyID")
			for _, occupied := range occupieds[index] {
				d.occupied[index][occupied.FamilyID] = append(d.occupied[index][occupied.FamilyID], storage.Reside{HouseID: occupied.HouseID, FamilyID: occupied.FamilyID, Level: occupied.Level, CheckOut: false})
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
				for index, occupied := range d.occupied[deleted.Level][deleted.FamilyID] {
					if occupied.HouseID == deleted.HouseID && occupied.CheckOut != true {
						d.occupied[deleted.Level][deleted.FamilyID], _ = storage.RemoveReside(d.occupied[deleted.Level][deleted.FamilyID], index)
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
		ctx.Respond(&distributorMessages.HouseCheckOutACK{})
		for _, occupied := range d.occupied[msg.Level][msg.FamilyID] {
			if occupied.CheckOut != true {
				occupied.CheckOut = true
				d.vacant[msg.Level] = append(d.vacant[msg.Level], storage.Reside{HouseID: occupied.HouseID, Level: msg.Level})
			}
			future := ctx.RequestFuture(d.managerPID, &distributorMessages.HouseCheckOut{FamilyID: occupied.FamilyID, HouseID: occupied.HouseID}, 2000*time.Millisecond)
			ctx.AwaitFuture(future, func(res interface{}, err error) {
				if err != nil {
					ctx.Self().Tell(msg)
					return
				}

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
		if msg.Retry != true {
			ctx.Respond(&distributorMessages.HouseApplicationACK{})
		}
		for _, occupied := range d.occupied[msg.Level][msg.FamilyID] {
			if occupied.CheckOut != true {
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
								d.occupied[msg.Level][msg.FamilyID], occupied = storage.RemoveReside(d.occupied[msg.Level][msg.FamilyID], i)
								break
							}
						}
						switch recv.Reason {
						case HAVEONEHOUSE, FAMILYDONOTEXIST:
							occupied.FamilyID = 0
							d.vacant[msg.Level] = append(d.vacant[msg.Level], occupied)
						case HOUSEMATCHED, HOUSEDONOTEXIST:
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
		ctx.Self().Tell(&distributorMessages.MatchEmptyHouse{Request: msg})
	case *distributorMessages.MatchEmptyHouse:
		if len(d.vacant[msg.Request.Level]) != 0 {
			var vacant storage.Reside
			if msg.Request.Retry != true {
				d.vacant[msg.Request.Level], vacant = storage.RemoveReside(d.vacant[msg.Request.Level], len(d.vacant[msg.Request.Level])-1)
				vacant.FamilyID = msg.Request.FamilyID
				d.occupied[msg.Request.Level][msg.Request.FamilyID] = append(d.occupied[msg.Request.Level][msg.Request.FamilyID], vacant)
			} else {
				vacant = d.occupied[msg.Request.Level][msg.Request.FamilyID][len(d.occupied[msg.Request.Level][msg.Request.FamilyID])-1]
			}
			// 发给 manager
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
					for i, occupied := range d.occupied[msg.Request.Level][msg.Request.FamilyID] {
						if occupied.CheckOut != true {
							d.occupied[msg.Request.Level][msg.Request.FamilyID], occupied = storage.RemoveReside(d.occupied[msg.Request.Level][msg.Request.FamilyID], i)
							break
						}
					}
					switch recv.Reason {
					case HAVEONEHOUSE, FAMILYDONOTEXIST:
						occupied.FamilyID = 0
						d.vacant[msg.Request.Level] = append(d.vacant[msg.Request.Level], occupied)
					case HOUSEMATCHED, HOUSEDONOTEXIST:
						msg.Request.Retry = false
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
