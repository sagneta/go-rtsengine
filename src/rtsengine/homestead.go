package rtsengine

/*
 Implements the  unit

*/

// HomeStead is an IUnit that maintains a homestead that generates peasants
type HomeStead struct {
	BaseUnit
}

func (unit *HomeStead) name() string {
	return "HomeStead"
}

func (unit *HomeStead) unitType() UnitType {
	return UnitHomeStead
}
