package rtsengine

/*
 Implements the farm unit

*/

// Tower is an IUnit that maintains a stone quarry
type Tower struct {
	Poolable
}

func (farm *Tower) name() string {
	return "Tower"
}
