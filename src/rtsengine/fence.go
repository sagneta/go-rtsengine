package rtsengine

/*
 Implements the unit

*/

// Fence is an IUnit that maintains a wood fence
type Fence struct {
	Poolable
	Health
	owner IPlayer
}

func (unit *Fence) name() string {
	return "Fence"
}
