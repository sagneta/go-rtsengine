package rtsengine

import (
	"math"
	"time"
)

/*
 Implements the  unit

*/

// Armory is an IUnit that maintains an Armory that produces combat units
type Armory struct {
	BaseUnit
}

func (unit *Armory) name() string {
	return "Armory"
}

func (unit *Armory) unitType() UnitType {
	return UnitArmory
}

func (unit *Armory) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 0
	unit.Life = 2000
	unit.AttackPoints = 1
	unit.AttackRange = 1
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = math.MaxInt64

	return unit
}
