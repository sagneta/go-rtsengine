package rtsengine

/*
 Implements the  unit

*/

// Goldmine is an IUnit that maintains a gold mine
type Goldmine struct {
	Poolable
	HealthAndAttack
	owner IPlayer
}

func (unit *Goldmine) name() string {
	return "Goldmine"
}

func (unit *Goldmine) unitType() UnitType {
	return UnitGoldMine
}
