package rtsengine

/*
 Implements the  unit

*/

// Goldmine is an IUnit that maintains a gold mine
type Goldmine struct {
	Poolable
	Health
	owner IPlayer
}

func (unit *Goldmine) name() string {
	return "Goldmine"
}
