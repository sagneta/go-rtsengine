package rtsengine

// UnitMap maintains a map of IUnits
type UnitMap struct {
	Map map[int]IUnit
}

// Remove will remove the unit from the map
func (unitmap *UnitMap) Remove(unit IUnit) {
	delete(unitmap.Map, unit.id()) // safe operation if unit not present.
}

// Add the unit to the UnitMap
func (unitmap *UnitMap) Add(unit IUnit) {
	unitmap.Map[unit.id()] = unit
}

// AllUnits returns the map
func (unitmap *UnitMap) AllUnits() map[int]IUnit {
	return unitmap.Map
}

// AddAll units to the UniMap.
func (unitmap *UnitMap) AddAll(objects ...IUnit) {
	for _, object := range objects {
		unitmap.Add(object)
	}
}
