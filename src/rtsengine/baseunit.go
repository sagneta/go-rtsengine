package rtsengine

// BaseUnit composes all the structures necessary for any unit.
type BaseUnit struct {
	Poolable
	AutoNumber
	HealthAndAttack
	owner IPlayer
}

// Initialize will set the unit to the base state.
// Call only once per instantiation
func (unit *BaseUnit) Initialize() {
	unit.Deallocate()
	unit.AutoNumber.Initialize()
}
