package rtsengine

/*
 Implements the  unit

*/

// StoneQuarry is an IUnit that maintains a stone quarry
type StoneQuarry struct {
	BaseUnit
}

func (unit *StoneQuarry) name() string {
	return "StoneQuarry"
}

func (unit *StoneQuarry) unitType() UnitType {
	return UnitStoneQuarry
}
