package rtsengine

/*
 Implements the  unit

*/

// Castle is an IUnit that maintains a military fortification
type Castle struct {
	BaseUnit
}

func (unit *Castle) name() string {
	return "Castle"
}

func (unit *Castle) unitType() UnitType {
	return UnitCastle
}
