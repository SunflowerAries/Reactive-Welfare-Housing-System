package property

import (
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/propertyMessages"
	"Reactive-Welfare-Housing-System/src/storage"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type propertyActor struct {
	db         storage.HouseSystem
	managerPID *actor.PID
}

// Exam the house, we need define a metric to do with
func ExamHouse(h storage.House) bool {
	return h.Age <= 5
}

func (p *propertyActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *managerMessages.ExaminationList:
		ExaminedHouse := p.db.QueryHouse(msg.HouseID)
		var houses []int32
		for _, house := range ExaminedHouse {
			if ExamHouse(house) != true {
				houses = append(houses, house.ID)
			}
		}
		if len(houses) == 0 {
			ctx.Respond(&propertyMessages.ExaminationACK{})
		} else {
			ctx.Respond(&propertyMessages.ExaminationRejects{HouseID: houses})
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
