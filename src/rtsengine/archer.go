package rtsengine

/*
 Implements the unit

*/

// Archer is an IUnit that maintains a range unit of some sort
type Archer struct {
	Poolable
}

func (unit *Archer) name() string {
	return "Archer"
}
