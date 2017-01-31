package rtsengine

/*
 Implements the farm unit

*/

// Peasant is an IUnit.
type Peasant struct {
	Poolable
}

func (farm *Peasant) name() string {
	return "Peasant"
}
