package rtsengine

/*
 Implements the unit

*/

// Cavalry is an IUnit that maintains a horse unit
type Cavalry struct {
	Poolable
}

func (unit *Cavalry) name() string {
	return "Cavalry"
}
