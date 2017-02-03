package rtsengine

/*
 Implements the  unit

*/

// Wall is an IUnit that maintains a stone (masonry) wall or defensive
// fortification like an abatis.
type Wall struct {
	Poolable
	HealthAndAttack
	owner IPlayer
}

func (unit *Wall) name() string {
	return "Wall"
}

func (unit *Wall) unitType() UnitType {
	return UnitWall
}
