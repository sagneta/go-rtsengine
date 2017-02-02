package rtsengine

/*
 Implements the unit

*/

// Cavalry is an IUnit that maintains a horse unit
type Cavalry struct {
	Poolable
	HealthAndAttack
	Movement
	owner IPlayer
}

func (unit *Cavalry) name() string {
	return "Cavalry"
}
