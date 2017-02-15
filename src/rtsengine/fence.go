package rtsengine

import (
	"math"
	"time"
)

/*
 Implements the unit

*/

// Fence is an IUnit that maintains a wood fence
type Fence struct {
	BaseUnit
}

func (unit *Fence) name() string {
	return "Fence"
}

func (unit *Fence) unitType() UnitType {
	return UnitFence
}

func (unit *Fence) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 1000
	unit.Life = 1000
	unit.AttackPoints = 0
	unit.AttackRange = 0
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = math.MaxInt64

	return unit
}
