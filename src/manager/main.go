package manager

import (
	"housingSystem/src/storage"

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
	case *supplyHouses:
		m.db.BatchInsertHouse(msg.houses)
	}
}

func NewManagerActor(db storage.HouseSystem) actor.Producer {
	return func() actor.Actor {
		return &manager{
			db: db,
		}
	}
}
