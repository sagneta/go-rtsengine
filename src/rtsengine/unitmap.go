package rtsengine

// UnitMap maintains a map of IUnits
type UnitMap struct {
	Map map[IUnit]IUnit
}

// Remove will remove the unit from the map
func (unitmap *UnitMap) Remove(unit IUnit) {
	delete(unitmap.Map, unit) // safe operation if unit not present.
}

// Add the unit to the UnitMap
func (unitmap *UnitMap) Add(unit IUnit) {
	unitmap.Map[unit] = unit
}

// AllUnits returns the map
func (unitmap *UnitMap) AllUnits() map[IUnit]IUnit {
	return unitmap.Map
}
