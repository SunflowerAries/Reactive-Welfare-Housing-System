package manager

import (
	"Reactive-Welfare-Housing-System/src/messages/distributorMessages"
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/storage"
	"fmt"
	"reflect"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type managerActor struct {
	db             storage.HouseSystem
	distributorPID *actor.PID
}

func (m *managerActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *sharedMessages.Houses:
		houses := make([]storage.House, 0, len(msg.Houses))
		for _, house := range msg.Houses {
			houses = append(houses, storage.House{Level: house.Level, Age: house.Age, Area: house.Area})
		}
		fmt.Println(houses, reflect.TypeOf(houses))
		m.db.BatchInsertHouse(houses)
	case *distributorMessages.HouseMatch:
		err := m.db.InsertMatch(storage.Reside{HouseID: msg.HouseID, FamilyID: msg.FamilyID})
		if err != nil {
			fmt.Printf("Insert house failed, err:%v", err)
			return
		}
		ctx.Respond(&managerMessages.HouseMatchApprove{Match: msg})
	case *sharedMessages.DistributorConnect:
		m.distributorPID = msg.Sender
		fmt.Println("In manager", msg.Sender)
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
