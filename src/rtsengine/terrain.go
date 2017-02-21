package rtsengine

/*
  Terrain enumeration.

*/

// Terrain enumeration
type Terrain byte

const (
	// Grass is the default.
	Grass Terrain = iota

	// Mountains are assumed to be any stoney structure
	Mountains

	// Trees are 0..N number of trees.
	Trees

	// Water is just that. An ocean is just a large number of these.
	Water

	// Sand is more or less a desert terrain.
	Sand

	// Snow is the white stuff.
	Snow

	// Dirt is dirt.
	Dirt
)
