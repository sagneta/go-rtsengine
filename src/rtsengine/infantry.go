package rtsengine

/*
 Implements the farm unit

*/

// Infantry is an IUnit that maintains a stone quarry
type Infantry struct {
	Poolable
}

func (farm *Infantry) name() string {
	return "Infantry"
}
