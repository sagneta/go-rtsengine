package rtsengine

import (
	"fmt"
	"sync"
)

/*
Each node in the world is an acre. It will contain numberous
elements such as a Unit structure and Terrain structure to name
just two.

*/

// Acre maintains the state for an acre of the World.
type Acre struct {
	terrain Terrain
	unit    IUnit

	// Synchronizes setting of a unit into this acre
	muSet sync.Mutex
}

// Set will set the newUnit into the acrea.
// Returns an error if there is already a unit
// occupying this acre.
func (acre *Acre) Set(newUnit IUnit) error {
	acre.muSet.Lock()
	defer acre.muSet.Unlock()
	if acre.unit != nil {
		return fmt.Errorf("Unit collision in acre")
	}
	acre.unit = newUnit
	return nil
}
