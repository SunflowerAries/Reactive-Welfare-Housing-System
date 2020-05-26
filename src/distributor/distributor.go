package distributor

import (
	"Reactive-Welfare-Housing-System/src/messages/distributorMessages"
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/messages/verifierMessages"
	"Reactive-Welfare-Housing-System/src/storage"
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type distributorActor struct {
	db          storage.HouseSystem
	occupied    [storage.HouseLevel + 1][]storage.Reside
	vacant      [storage.HouseLevel + 1][]storage.Reside
	allocating  [storage.HouseLevel + 1][]storage.Reside
	managerPID  *actor.PID
	verifierPID *actor.PID
}

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
	case *managerMessages.HouseMatchApprove:
		var allocating storage.Reside
		d.allocating[msg.Match.Level], allocating = storage.RemoveReside(d.allocating[msg.Match.Level], int(msg.Match.Index))
		allocating.CheckOut = false
		if allocating.FamilyID != msg.Match.FamilyID || allocating.HouseID != msg.Match.HouseID {
			panic("msg not equal memory database")
		}
		d.occupied[allocating.Level] = append(d.allocating[allocating.Level], allocating)
		request := &verifierMessages.HouseApplicationRequest{FamilyID: msg.Match.FamilyID, Level: msg.Match.Level}
		ctx.Request(d.verifierPID, &distributorMessages.HouseApplicationResponse{Request: request, HouseID: msg.Match.HouseID})
		// 发给 tenant?
	case *managerMessages.HouseMatchReject:
		if len(d.vacant[msg.Match.Level]) != 0 {
			var vacant storage.Reside
			d.allocating[msg.Match.Level], _ = storage.RemoveReside(d.allocating[msg.Match.Level], int(msg.Match.Index))
			d.vacant[msg.Match.Level], vacant = storage.RemoveReside(d.vacant[msg.Match.Level], len(d.vacant[msg.Match.Level])-1)
			vacant.FamilyID = msg.Match.FamilyID
			vacant.Level = msg.Match.Level
			d.allocating[msg.Match.Level] = append(d.allocating[msg.Match.Level], vacant)
			// 发给 manager
			ctx.Request(d.managerPID, &distributorMessages.HouseMatch{FamilyID: vacant.FamilyID, HouseID: vacant.HouseID, Index: int32(len(d.allocating[msg.Match.Level]) - 1)})
			return
		}
		// 发给 tenant
		request := &verifierMessages.HouseApplicationRequest{FamilyID: msg.Match.FamilyID, Level: msg.Match.Level}
		ctx.Request(d.verifierPID, &distributorMessages.HouseApplicationReject{Request: request, Reason: "对不起，系统中暂时无匹配住房"})
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
						if occupied.CheckOut != true {
							ctx.Request(d.verifierPID, &distributorMessages.HouseRecall{FamilyID: deleted.FamilyID, HouseID: deleted.House.ID})
						}
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
	case *verifierMessages.HouseCheckOut:
		for _, occupied := range d.occupied[msg.Level] {
			if occupied.FamilyID == msg.FamilyID {
				if occupied.CheckOut != false {
					// 发给 tenant?
					ctx.Respond(&distributorMessages.HouseCheckOutResponse{Answer: "退房请求正在处理中"})
					return
				}
				occupied.CheckOut = true
				ctx.Request(d.managerPID, &distributorMessages.HouseCheckOut{FamilyID: msg.FamilyID, HouseID: occupied.HouseID})
				d.vacant[msg.Level] = append(d.vacant[msg.Level], storage.Reside{HouseID: occupied.HouseID, Level: msg.Level})
				return
			}
		}
		for _, allocating := range d.allocating[msg.Level] {
			if allocating.FamilyID == msg.FamilyID {
				// 发给 tenant
				ctx.Respond(&distributorMessages.HouseCheckOutResponse{Answer: "租房请求正在处理中"})
				return
			}
		}
	case *verifierMessages.HouseApplicationRequest:
		for _, occupied := range d.occupied[msg.Level] {
			if occupied.FamilyID == msg.FamilyID && occupied.CheckOut != true {
				// 发给 tenant
				ctx.Respond(&distributorMessages.HouseApplicationReject{Request: msg, Reason: "对不起，一户家庭不能租住超过一套保障房"})
				return
			}
		}
		for _, allocating := range d.allocating[msg.Level] {
			if allocating.FamilyID == msg.FamilyID {
				// 发给 tenant
				ctx.Respond(&distributorMessages.HouseApplicationReject{Request: msg, Reason: "对不起，正在为您分配住房，请耐心等待"})
				return
			}
		}
		if len(d.vacant[msg.Level]) != 0 {
			var vacant storage.Reside
			d.vacant[msg.Level], vacant = storage.RemoveReside(d.vacant[msg.Level], len(d.vacant[msg.Level])-1)
			vacant.FamilyID = msg.FamilyID
			vacant.Level = msg.Level
			d.allocating[msg.Level] = append(d.allocating[msg.Level], vacant)
			// 发给 manager
			ctx.Request(d.managerPID, &distributorMessages.HouseMatch{FamilyID: vacant.FamilyID, HouseID: vacant.HouseID, Index: int32(len(d.allocating[msg.Level]) - 1)})
			return
		}
		// 发给 tenant
		ctx.Respond(&distributorMessages.HouseApplicationReject{Request: msg, Reason: "对不起，系统中暂时无匹配住房"})
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
