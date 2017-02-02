package rtsengine

/*
 Implements the  unit

*/

// HomeStead is an IUnit that maintains a homestead that generates peasants
type HomeStead struct {
	Poolable
	HealthAndAttack
	owner IPlayer
}

func (unit *HomeStead) name() string {
	return "HomeStead"
}
