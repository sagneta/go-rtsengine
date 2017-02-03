package rtsengine

/*
 Implements the Catapult (artillery) unit

*/

// Catapult is an IUnit that maintains artillery of some sort.
type Catapult struct {
	Poolable
	HealthAndAttack
	Movement
	owner IPlayer
}

func (unit *Catapult) name() string {
	return "Catapult"
}

func (unit *Catapult) unitType() UnitType {
	return UnitCatapult
}
