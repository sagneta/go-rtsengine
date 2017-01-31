package rtsengine

/*
 Implements the farm unit

*/

// Cavalry is an IUnit that maintains a stone quarry
type Cavalry struct {
	Poolable
}

func (farm *Cavalry) name() string {
	return "Cavalry"
}
