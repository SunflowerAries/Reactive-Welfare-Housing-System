package manager

import (
	"Reactive-Welfare-Housing-System/src/messages/distributorMessages"
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/propertyMessages"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/property"
	"Reactive-Welfare-Housing-System/src/storage"
	"fmt"
	"reflect"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/router"
)

type managerActor struct {
	db             storage.HouseSystem
	distributorPID *actor.PID
	propertyPID    *actor.PID
}

func (m *managerActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		m.propertyPID = ctx.Spawn(router.NewRoundRobinPool(5).WithProducer(property.NewPropertyActor(m.db, ctx.Self())))
	case *sharedMessages.NewHouses:
		var houses []storage.House
		for _, house := range msg.Houses {
			houses = append(houses, storage.House{Level: house.Level, Age: house.Age, Area: house.Area})
		}
		HouseID := m.db.BatchInsertHouse(houses)
		start := 0
		for i := 0; i < len(HouseID)/2; i++ {
			for j := 0; j < HouseID[2*i]; j++ {
				houses[start+j].ID = int32(HouseID[2*i+1] + j)
			}
			start += HouseID[2*i]
		}
		newhouses := make([]*sharedMessages.NewHouse, 0, len(msg.Houses))
		for _, house := range houses {
			newhouses = append(newhouses, &sharedMessages.NewHouse{ID: house.ID, Level: house.Level})
		}
		ctx.Send(m.distributorPID, &managerMessages.NewHouses{Houses: &sharedMessages.NewHouses{Houses: newhouses}})
	// TCP
	case *sharedMessages.DistributorConnect:
		m.distributorPID = msg.Sender
		fmt.Println("In manager", msg.Sender)
	case *sharedMessages.ExaminationList:
		ctx.Request(m.propertyPID, &managerMessages.ExaminationList{HouseID: msg.HouseID})
	case *distributorMessages.HouseMatch:
		err := m.db.InsertMatch(storage.Reside{HouseID: msg.HouseID, FamilyID: msg.FamilyID})
		if err != nil {
			fmt.Printf("Insert house failed, err:%v", err)
			ctx.Respond(&managerMessages.HouseMatchReject{Match: msg})
			return
		}
		ctx.Respond(&managerMessages.HouseMatchApprove{Match: msg})
	case *distributorMessages.HouseCheckOut:
		err := m.db.CheckOutHouse(storage.Reside{FamilyID: msg.FamilyID, HouseID: msg.HouseID})
		if err != nil {
			fmt.Printf("Checkout house failed, err:%v", err)
			return
		}
	case *propertyMessages.ExaminationRejects:
		var houses []*propertyMessages.ExaminationReject
		var ids []int32
		for _, house := range msg.Houses {
			houses = append(houses, house)
			ids = append(ids, house.House.ID)
		}
		err := m.db.DeleteHouse(ids)
		if err != nil {
			fmt.Printf("Delete house failed, err:%v", err)
			return
		}
		ctx.Request(m.distributorPID, &managerMessages.UnqualifiedHouses{Houses: houses})
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
