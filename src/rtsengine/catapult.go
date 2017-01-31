package rtsengine

/*
 Implements the Catapult (artillery) unit

*/

// Catapult is an IUnit that maintains artillery of some sort.
type Catapult struct {
	Poolable
	Health
	owner IPlayer
}

func (unit *Catapult) name() string {
	return "Catapult"
}
