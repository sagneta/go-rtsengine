package rtsengine

/*
 Implements the farm unit

*/

// Catapult is an IUnit that maintains a stone quarry
type Catapult struct {
	Poolable
}

func (farm *Catapult) name() string {
	return "Catapult"
}
