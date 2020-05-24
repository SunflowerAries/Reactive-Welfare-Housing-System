package government

import (
	"housingSystem/src/storage"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type government struct{
	db 	storage.HouseSystem
}




func (t *government) Receive(ctx actor.Context) {
	switch msg : ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("Government Starting, initial actor here, PID:", ctx.Self())
		t.self = ctx.Self()
	case *actor.Stopped:
		fmt.Println("Government Stopped")
	}
}

func newGovernmentActor() actor.Actor {
	return &Actor{}
}