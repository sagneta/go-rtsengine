package rtsengine

/*
 Implements the farm unit

*/

// Goldmine is an IUnit that maintains a gold mine
type Goldmine struct {
	Poolable
}

func (farm *Goldmine) name() string {
	return "Goldmine"
}
