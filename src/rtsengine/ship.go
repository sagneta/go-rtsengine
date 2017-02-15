package rtsengine

import "time"

/*
 Implements the  unit

*/

// Ship is an IUnit that maintains a military vessel of some sort.
type Ship struct {
	BaseUnit
}

func (unit *Ship) name() string {
	return "Ship"
}

func (unit *Ship) unitType() UnitType {
	return UnitShip
}

func (unit *Ship) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 100
	unit.Life = 100
	unit.AttackPoints = 50
	unit.AttackRange = 20
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = 1000

	return unit
}
