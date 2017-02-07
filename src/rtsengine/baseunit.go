package rtsengine

import "math"

// BaseUnit composes all the structures necessary for any unit.
type BaseUnit struct {
	Poolable
	AutoNumber
	HealthAndAttack
	owner IPlayer
	Movement
}

// Initialize will set the unit to the base state.
// Call only once per instantiation
func (unit *BaseUnit) Initialize() {
	unit.Deallocate()
	unit.AutoNumber.Initialize()

	// Default to something huge which make it immovable.
	// Override if you want the unit to move.
	unit.DeltaInMillis = math.MaxInt64
}

// IUnit
func (unit *BaseUnit) movement() *Movement {
	return &unit.Movement
}
