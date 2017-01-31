package rtsengine

/*
 Implements the farm unit

*/

// Ship is an IUnit that maintains a stone quarry
type Ship struct {
	Poolable
}

func (farm *Ship) name() string {
	return "Ship"
}
