package rtsengine

import "image"

/*
 World 2D grid. That is an array of acre structures.

*/

// World maintains the world state. This is the big one!
type World struct {
	Grid [][]Acre
	span image.Rectangle
}

// NewWorld will construct a random world of width and height specified.
// works on 'this'. Another way of thinking is width are the columns
// and height are the rows.
func NewWorld(width int, height int) *World {
	world := World{}

	// allocate 2d array row per row.
	world.Grid = make([][]Acre, height)
	for i := range world.Grid {
		world.Grid[i] = make([]Acre, width)
	}

	// store the dimensions for later.
	world.span = image.Rect(0, 0, width, height)

	// Generate the entire world semi-randomly
	// We will need some configuration parameters
	// to control this behavior.
	world.Grid[0][0].terrain = Trees

	return &world
}
