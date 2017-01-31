package rtsengine

/*
 Implements the  unit

*/

// Wall is an IUnit that maintains a stone (masonry) wall or defensive
// fortification like an abatis.
type Wall struct {
	Poolable
	Health
	owner IPlayer
}

func (unit *Wall) name() string {
	return "Wall"
}
