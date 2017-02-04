package rtsengine

/*
 Implements the unit

*/

// Archer is an IUnit that maintains a range unit of some sort
type Archer struct {
	BaseUnit
	Movement
}

func (unit *Archer) name() string {
	return "Archer"
}

func (unit *Archer) unitType() UnitType {
	return UnitArcher
}
