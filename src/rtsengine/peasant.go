package rtsengine

/*
 Implements the  unit

*/

// Peasant is an IUnit. Basic non-combatant that produces resources.
type Peasant struct {
	Poolable
	HealthAndAttack
	Movement
	owner IPlayer
}

func (unit *Peasant) name() string {
	return "Peasant"
}
