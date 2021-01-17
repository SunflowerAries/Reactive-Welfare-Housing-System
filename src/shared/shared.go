package shared

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

// Use use the variable to avoid <declaration but not used> err.
func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

// BaseActor The Base Actor
type BaseActor struct {
	Self *actor.PID
	Name string
}
