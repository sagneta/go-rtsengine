package rtsengine

/*
Each node in the world is an acre. It will contain numberous
elements such as a Unit structure and Terrain structure to name
just two.

*/

// Acre maintains the state for an acre of the World.
type Acre struct {
	terrain Terrain
	unit    IUnit
}
