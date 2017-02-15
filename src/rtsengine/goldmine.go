package rtsengine

import (
	"math"
	"time"
)

/*
 Implements the  unit

*/

// Goldmine is an IUnit that maintains a gold mine
type Goldmine struct {
	BaseUnit
}

func (unit *Goldmine) name() string {
	return "Goldmine"
}

func (unit *Goldmine) unitType() UnitType {
	return UnitGoldMine
}

func (unit *Goldmine) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 100000
	unit.Life = 100000
	unit.AttackPoints = 0
	unit.AttackRange = 0
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = math.MaxInt64

	return unit
}
