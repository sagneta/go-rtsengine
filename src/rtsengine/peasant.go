package rtsengine

import (
	"math"
	"time"
)

/*
 Implements the  unit

*/

// Peasant is an IUnit. Basic non-combatant that produces resources.
type Peasant struct {
	BaseUnit
}

func (unit *Peasant) name() string {
	return "Peasant"
}

func (unit *Peasant) unitType() UnitType {
	return UnitPeasant
}

func (unit *Peasant) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 10
	unit.Life = 10
	unit.AttackPoints = 0
	unit.AttackRange = 0
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = math.MaxInt64

	return unit
}
