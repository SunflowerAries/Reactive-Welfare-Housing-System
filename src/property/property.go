package property

import (
	"Reactive-Welfare-Housing-System/src/messages/managerMessages"
	"Reactive-Welfare-Housing-System/src/messages/propertyMessages"
	"Reactive-Welfare-Housing-System/src/storage"
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type propertyActor struct {
	db         storage.HouseSystem
	managerPID *actor.PID
}

// Exam the house, we need define a metric to do with
func ExamHouse(h storage.Reside) bool {
	return h.Age <= 5
}

func (p *propertyActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *managerMessages.ExaminationList:
		ExaminedHouse := p.db.QueryHouse(msg.HouseID)
		var houses []*propertyMessages.ExaminationReject
		for _, house := range ExaminedHouse {
			fmt.Println(house)
			if ExamHouse(house) != true {
				rejected := &propertyMessages.ExaminedHouse{ID: house.HouseID, Level: house.Level, Age: house.Age, Area: house.Area}
				houses = append(houses, &propertyMessages.ExaminationReject{House: rejected, FamilyID: house.FamilyID})
			}
		}
		if len(houses) == 0 {
			ctx.Respond(&propertyMessages.ExaminationACK{})
		} else {
			ctx.Respond(&propertyMessages.ExaminationRejects{Houses: houses})
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
