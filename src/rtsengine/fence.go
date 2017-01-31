package rtsengine

/*
 Implements the farm unit

*/

// Fence is an IUnit that maintains a stone quarry
type Fence struct {
	Poolable
}

func (farm *Fence) name() string {
	return "Fence"
}
