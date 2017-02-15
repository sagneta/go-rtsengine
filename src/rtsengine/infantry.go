package rtsengine

import "time"

/*
 Implements the  unit

*/

// Infantry is an IUnit that maintains a foot soldier unit. Like a century or company
type Infantry struct {
	BaseUnit
}

func (unit *Infantry) name() string {
	return "Infantry"
}

func (unit *Infantry) unitType() UnitType {
	return UnitInfantry
}

func (unit *Infantry) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 100
	unit.Life = 100
	unit.AttackPoints = 2
	unit.AttackRange = 1
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = 1000

	return unit
}
