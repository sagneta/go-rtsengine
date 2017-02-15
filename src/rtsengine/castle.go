package rtsengine

import (
	"math"
	"time"
)

/*
 Implements the  unit

*/

// Castle is an IUnit that maintains a military fortification
type Castle struct {
	BaseUnit
}

func (unit *Castle) name() string {
	return "Castle"
}

func (unit *Castle) unitType() UnitType {
	return UnitCastle
}

func (unit *Castle) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 5000
	unit.Life = 5000
	unit.AttackPoints = 10
	unit.AttackRange = 5
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = math.MaxInt64

	return unit
}
