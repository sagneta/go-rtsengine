package rtsengine

/*
 Implements the  unit

*/

// Ship is an IUnit that maintains a military vessel of some sort.
type Ship struct {
	Poolable
	Health
	owner IPlayer
}

func (unit *Ship) name() string {
	return "Ship"
}
