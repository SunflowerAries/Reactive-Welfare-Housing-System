package shared

import (
	"github.com/ChaokunChang/protoactor-go/actor"
)

const HAVEONEHOUSE = 1
const HOUSEMATCHED = 2
const HOUSEDONOTEXIST = 3
const FAMILYDONOTEXIST = 4

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
