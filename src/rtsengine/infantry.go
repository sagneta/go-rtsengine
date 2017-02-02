package rtsengine

/*
 Implements the  unit

*/

// Infantry is an IUnit that maintains a foot soldier unit. Like a century or company
type Infantry struct {
	Poolable
	HealthAndAttack
	Movement
	owner IPlayer
}

func (unit *Infantry) name() string {
	return "Infantry"
}
