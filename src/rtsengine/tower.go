package rtsengine

import (
	"math"
	"time"
)

/*
 Implements the  unit

*/

// Tower is an IUnit that maintains a military watch tower with some defensive capability
type Tower struct {
	BaseUnit
}

func (unit *Tower) name() string {
	return "Tower"
}

func (unit *Tower) unitType() UnitType {
	return UnitTower
}

func (unit *Tower) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 1000
	unit.Life = 1000
	unit.AttackPoints = 0
	unit.AttackRange = 0
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = math.MaxInt64

	return unit
}
