package rtsengine

/*
 Implements the unit

*/

// Archer is an IUnit that maintains a range unit of some sort
type Archer struct {
	Poolable
	HealthAndAttack
	Movement
	owner IPlayer
}

func (unit *Archer) name() string {
	return "Archer"
}

func (unit *Archer) unitType() UnitType {
	return UnitArcher
}
