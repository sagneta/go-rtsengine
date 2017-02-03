package rtsengine

/*
 Implements the  unit

*/

// Castle is an IUnit that maintains a military fortification
type Castle struct {
	Poolable
	HealthAndAttack
	owner IPlayer
}

func (unit *Castle) name() string {
	return "Castle"
}

func (unit *Castle) unitType() UnitType {
	return UnitCastle
}
