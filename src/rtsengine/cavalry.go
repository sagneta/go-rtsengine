package rtsengine

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
