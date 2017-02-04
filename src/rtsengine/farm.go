package rtsengine

/*
 Implements the farm unit

*/

// Farm is an IUnit that maintains a farm and adds food resources to an IPlayer
type Farm struct {
	BaseUnit
}

func (farm *Farm) name() string {
	return "Farm"
}

func (farm *Farm) unitType() UnitType {
	return UnitFarm
}
