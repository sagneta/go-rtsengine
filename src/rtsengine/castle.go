package rtsengine

/*
 Implements the farm unit

*/

// Castle is an IUnit that maintains a stone quarry
type Castle struct {
	Poolable
}

func (farm *Castle) name() string {
	return "Castle"
}
