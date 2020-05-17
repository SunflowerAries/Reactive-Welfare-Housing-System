package main

import (
	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/jinzhu/gorm"
)

type applicant struct{}

type distributor struct {
	MysqlDB *gorm.DB
}

func (d *distributor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *applicant:
		result := d.match(*msg)
		ctx.Send(ctx.Self(), result)
	}
}

func (d *distributor) match(p applicant) {
	// look up db to match
	return "asds"
}

var context *actor.RootContext

func main() {
	context = actor.EmptyRootContext
	distributorProps := actor.PropsFromProducer(func() actor.Actor { return &distributor{} })
	context.Spawn(distributorProps)

	console.ReadLine()
}
