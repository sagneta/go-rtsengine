package rtsengine

/*
 Implements the farm unit

*/

// StoneQuarry is an IUnit that maintains a stone quarry
type StoneQuarry struct {
	Poolable
}

func (farm *StoneQuarry) name() string {
	return "StoneQuarry"
}
