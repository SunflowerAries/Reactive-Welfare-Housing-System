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
	db         storage.HouseSystem
	occupied   [storage.HouseLevel + 1][]storage.Reside
	vacant     [storage.HouseLevel + 1][]storage.Reside
	allocating [storage.HouseLevel + 1][]storage.Reside
	managerPID *actor.PID
}

func (d *distributorActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		d.occupied, d.vacant = d.db.InitMatchCache()
		fmt.Printf("%+v\t%+v\t%+v\n", d.occupied, d.vacant, msg)
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

	case *verifierMessages.HouseApplicationRequest:
		for _, occupied := range d.occupied[msg.Level] {
			if occupied.FamilyID == msg.FamilyID && occupied.CheckOut != true {
				// 发给 tenant
				ctx.Respond(&distributorMessages.HouseApplicationReject{Reason: "对不起，一户家庭不能租住超过一套保障房"})
				return
			}
		}
		for _, allocating := range d.allocating[msg.Level] {
			if allocating.FamilyID == msg.FamilyID {
				// 发给 tenant
				ctx.Respond(&distributorMessages.HouseApplicationReject{Reason: "对不起，正在为您分配住房，请耐心等待"})
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
		} else {
			// 发给 tenant
			ctx.Respond(&distributorMessages.HouseApplicationReject{Reason: "对不起，系统中暂时无匹配住房"})
		}
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
