package rtsengine

import "time"

/*
 Implements the Catapult (artillery) unit

*/

// Catapult is an IUnit that maintains artillery of some sort.
type Catapult struct {
	BaseUnit
	Movement
}

func (unit *Catapult) name() string {
	return "Catapult"
}

func (unit *Catapult) unitType() UnitType {
	return UnitCatapult
}

func (unit *Catapult) generate(player IPlayer) IUnit {
	unit.Owner = player
	unit.HitPoints = 100
	unit.Life = 100
	unit.AttackPoints = 50
	unit.AttackRange = 10
	unit.LastMovement = time.Now()
	unit.DeltaInMillis = 1000

	return unit
}
