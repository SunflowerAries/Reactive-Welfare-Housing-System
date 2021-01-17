package property

import (
	"Reactive-Welfare-Housing-System/src/messages/propertyMessages"
	"Reactive-Welfare-Housing-System/src/messages/sharedMessages"
	"Reactive-Welfare-Housing-System/src/storage"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type propertyActor struct {
	db         storage.HouseSystem
	managerPID *actor.PID
}

// Exam the house, we need define a metric to do with
func examHouse(h storage.House) bool {
	return h.Age <= 5 || h.Deleted == true
}

func (p *propertyActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *sharedMessages.ExaminationList:
		examinedHouses := p.db.QueryHouse(msg.HouseID)
		var houses []int32
		for _, house := range examinedHouses {
			if examHouse(house) != true {
				houses = append(houses, house.ID)
			}
		}
		if len(houses) != 0 {
			ctx.Request(p.managerPID, &propertyMessages.UnqualifiedHouseIDs{HouseID: houses})
		}
	}
}

func NewPropertyActor(db storage.HouseSystem, managerPID *actor.PID) actor.Producer {
	return func() actor.Actor {
		return &propertyActor{
			db:         db,
			managerPID: managerPID,
		}
	}
}
