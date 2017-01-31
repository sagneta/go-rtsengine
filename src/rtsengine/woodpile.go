package rtsengine

/*
 Implements the farm unit

*/

// WoodPile is an IUnit that maintains a stone quarry
type WoodPile struct {
	Poolable
}

func (farm *WoodPile) name() string {
	return "WoodPile"
}
