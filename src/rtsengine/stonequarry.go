package rtsengine

import (
	"math"
	"time"
)

/*
 Implements the  unit

*/

// StoneQuarry is an IUnit that maintains a stone quarry
type StoneQuarry struct {
	BaseUnit
}

func (unit *StoneQuarry) name() string {
	return "StoneQuarry"
}

func (unit *StoneQuarry) unitType() UnitType {
	return UnitStoneQuarry
}

func (unit *StoneQuarry) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 100000
	unit.Life = 100000
	unit.AttackPoints = 0
	unit.AttackRange = 0
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = math.MaxInt64

	return unit
}
