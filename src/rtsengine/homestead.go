package rtsengine

import "time"

/*
 Implements the  unit

*/

// HomeStead is an IUnit that maintains a homestead that generates peasants
type HomeStead struct {
	BaseUnit
}

func (unit *HomeStead) name() string {
	return "HomeStead"
}

func (unit *HomeStead) unitType() UnitType {
	return UnitHomeStead
}

func (unit *HomeStead) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 1000
	unit.Life = 1000
	unit.AttackPoints = 1
	unit.AttackRange = 2
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = 1000

	return unit
}
