package rtsengine

/*
 Implements the  unit

*/

// StoneQuarry is an IUnit that maintains a stone quarry
type StoneQuarry struct {
	Poolable
	Health
	owner IPlayer
}

func (unit *StoneQuarry) name() string {
	return "StoneQuarry"
}
