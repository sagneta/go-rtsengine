package rtsengine

/*
 Implements the  unit

*/

// Ship is an IUnit that maintains a military vessel of some sort.
type Ship struct {
	BaseUnit
	Movement
}

func (unit *Ship) name() string {
	return "Ship"
}

func (unit *Ship) unitType() UnitType {
	return UnitShip
}
