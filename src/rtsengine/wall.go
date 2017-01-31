package rtsengine

/*
 Implements the farm unit

*/

// Wall is an IUnit that maintains a stone quarry
type Wall struct {
	Poolable
}

func (farm *Wall) name() string {
	return "Wall"
}
