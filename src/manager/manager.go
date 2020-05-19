package manager

import (
	"fmt"
	"housingSystem/src/messages"
	"housingSystem/src/storage"
	"reflect"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type supplyHouses struct {
	houses []storage.House
}

type manager struct {
	db storage.HouseSystem
}

func (m *manager) Receive(ctx actor.Context) {

	switch msg := ctx.Message().(type) {
	case *messages.Houses:
		fmt.Println("In supplyhouses")
		fmt.Printf("%+v\n", msg.Houses)
		houses := make([]storage.House, 0, len(msg.Houses))
		for _, house := range msg.Houses {
			houses = append(houses, storage.House{Level: house.Level, Age: house.Age, Area: house.Area})
		}
		fmt.Println(houses, reflect.TypeOf(houses))
		m.db.BatchInsertHouse(houses)
	default:
		fmt.Println(reflect.TypeOf(msg), msg)
	}
}

func NewManagerActor(db storage.HouseSystem) actor.Producer {
	return func() actor.Actor {
		return &manager{
			db: db,
		}
	}
}
