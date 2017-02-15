package rtsengine

import (
	"math"
	"time"
)

/*
 Implements the farm unit

*/

// Farm is an IUnit that maintains a farm and adds food resources to an IPlayer
type Farm struct {
	BaseUnit
}

func (unit *Farm) name() string {
	return "Farm"
}

func (unit *Farm) unitType() UnitType {
	return UnitFarm
}

func (unit *Farm) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 10
	unit.Life = 10
	unit.AttackPoints = 0
	unit.AttackRange = 0
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = math.MaxInt64

	return unit
}
