package rtsengine

/*
 Implements the  unit

*/

// StoneQuarry is an IUnit that maintains a stone quarry
type StoneQuarry struct {
	Poolable
}

func (unit *StoneQuarry) name() string {
	return "StoneQuarry"
}
