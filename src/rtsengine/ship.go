package rtsengine

/*
 Implements the  unit

*/

// Ship is an IUnit that maintains a military vessel of some sort.
type Ship struct {
	Poolable
	HealthAndAttack
	Movement
	owner IPlayer
}

func (unit *Ship) name() string {
	return "Ship"
}

func (unit *Ship) unitType() UnitType {
	return UnitShip
}
