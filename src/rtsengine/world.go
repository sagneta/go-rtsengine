package rtsengine

import "image"

/*
 World 2D grid. That is an array of acre structures.

*/

// World maintains the world state. This is the big one!
type World struct {
	Grid
}

// NewWorld will construct a random world of width and height specified.
// works on 'this'. Another way of thinking is width are the columns
// and height are the rows.
func NewWorld(width int, height int) *World {
	world := World{}

	// When the worldLocation is 0,0 then the grid IS the world.
	world.GenerateGrid(image.Point{0, 0}, width, height)

	// Generate the entire world semi-randomly
	// We will need some configuration parameters
	// to control this behavior.
	world.Matrix[0][0].terrain = Trees

	return &world
}
