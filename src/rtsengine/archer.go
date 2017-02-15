package rtsengine

import "time"

/*
 Implements the unit

*/

// Archer is an IUnit that maintains a range unit of some sort
type Archer struct {
	BaseUnit
}

func (unit *Archer) name() string {
	return "Archer"
}

func (unit *Archer) unitType() UnitType {
	return UnitArcher
}

func (unit *Archer) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 10
	unit.Life = 40
	unit.AttackPoints = 1
	unit.AttackRange = 5
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = 1000

	return unit
}
