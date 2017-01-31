package rtsengine

/*
 Implements the farm unit

*/

// Archer is an IUnit that maintains a stone quarry
type Archer struct {
	Poolable
}

func (farm *Archer) name() string {
	return "Archer"
}
