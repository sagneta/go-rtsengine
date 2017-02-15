package rtsengine

import (
	"math"
	"time"
)

/*
 Implements the  unit

*/

// WoodPile is an IUnit that maintains a wood pile that provides wood.
type WoodPile struct {
	BaseUnit
}

func (unit *WoodPile) name() string {
	return "WoodPile"
}

func (unit *WoodPile) unitType() UnitType {
	return UnitWoodPile
}

func (unit *WoodPile) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 100000
	unit.Life = 100000
	unit.AttackPoints = 0
	unit.AttackRange = 0
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = math.MaxInt64

	return unit
}
