package rtsengine

import "time"

/*
 Implements the unit

*/

// Cavalry is an IUnit that maintains a horse unit
type Cavalry struct {
	BaseUnit
}

func (unit *Cavalry) name() string {
	return "Cavalry"
}

func (unit *Cavalry) unitType() UnitType {
	return UnitCavalry
}

func (unit *Cavalry) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 100
	unit.Life = 100
	unit.AttackPoints = 4
	unit.AttackRange = 1
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = 400

	return unit
}
