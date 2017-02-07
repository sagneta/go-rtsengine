package rtsengine

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
