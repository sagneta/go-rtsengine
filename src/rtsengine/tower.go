package rtsengine

/*
 Implements the  unit

*/

// Tower is an IUnit that maintains a military watch tower with some defensive capability
type Tower struct {
	Poolable
	HealthAndAttack
	owner IPlayer
}

func (unit *Tower) name() string {
	return "Tower"
}

func (unit *Tower) unitType() UnitType {
	return UnitTower
}
