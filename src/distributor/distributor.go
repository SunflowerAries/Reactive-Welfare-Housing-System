package distributor

import (
	"fmt"
	"housingSystem/src/storage"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type applicant struct{}

type distributor struct {
	db       storage.HouseSystem
	occupied []storage.Reside
	vacant   [storage.HouseLevel + 1][]uint
}

func (d *distributor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		d.occupied, d.vacant = d.db.InitMatchCache()
		fmt.Printf("%+v\t%+v\t%+v\n", d.occupied, d.vacant, msg)
		// case *applicant:
		// 	// result := d.match(*msg)
		// 	ctx.Send(ctx.Self(), result)
		// }
	}
}

func NewDistributorActor(db storage.HouseSystem) actor.Producer {
	return func() actor.Actor {
		return &distributor{
			db: db,
		}
	}
}
