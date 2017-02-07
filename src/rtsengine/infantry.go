package rtsengine

/*
 Implements the  unit

*/

// Infantry is an IUnit that maintains a foot soldier unit. Like a century or company
type Infantry struct {
	BaseUnit
}

func (unit *Infantry) name() string {
	return "Infantry"
}

func (unit *Infantry) unitType() UnitType {
	return UnitInfantry
}
