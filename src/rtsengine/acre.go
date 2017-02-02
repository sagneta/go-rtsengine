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
	if acre.Collision() {
		return fmt.Errorf("Unit collision in acre.")
	}
	acre.unit = newUnit
	return nil
}

// Collision returns true if the locus is already occupied
// by any other unit OR the terrain is inaccessible such as
// Mountains, Water and Trees.
func (acre *Acre) Collision() bool {
	return acre.unit != nil || acre.terrain == Trees || acre.terrain == Mountains || acre.terrain == Water
}

// Occupied returns true if this acre is occupied by a unit and false otherwise.
func (acre *Acre) Occupied() bool {
	return acre.unit != nil
}
