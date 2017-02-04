package rtsengine

/*
 Implements the Catapult (artillery) unit

*/

// Catapult is an IUnit that maintains artillery of some sort.
type Catapult struct {
	BaseUnit
	Movement
}

func (unit *Catapult) name() string {
	return "Catapult"
}

func (unit *Catapult) unitType() UnitType {
	return UnitCatapult
}
