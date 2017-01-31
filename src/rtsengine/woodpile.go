package rtsengine

/*
 Implements the  unit

*/

// WoodPile is an IUnit that maintains a wood pile that provides wood.
type WoodPile struct {
	Poolable
}

func (unit *WoodPile) name() string {
	return "WoodPile"
}
