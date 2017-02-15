package rtsengine

import (
	"math"
	"time"
)

/*
 Implements the  unit

*/

// Wall is an IUnit that maintains a stone (masonry) wall or defensive
// fortification like an abatis.
type Wall struct {
	BaseUnit
}

func (unit *Wall) name() string {
	return "Wall"
}

func (unit *Wall) unitType() UnitType {
	return UnitWall
}

func (unit *Wall) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 3000
	unit.Life = 3000
	unit.AttackPoints = 0
	unit.AttackRange = 0
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = math.MaxInt64

	return unit
}
