package rtsengine

/*
 Implements the unit

*/

// Fence is an IUnit that maintains a wood fence
type Fence struct {
	BaseUnit
}

func (unit *Fence) name() string {
	return "Fence"
}

func (unit *Fence) unitType() UnitType {
	return UnitFence
}
