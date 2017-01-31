package rtsengine

import "image"

/*
 A View into the world grid

*/

// View is a projection onto the World Grid
type View struct {
	Grid
}

// GenerateView will initialize all internal structures.
// It will set the grid widith and height and situate the
// view onto the world at worldLocation
func (view *View) GenerateView(worldLocation image.Point, width int, height int) {
	view.GenerateGrid(worldLocation, width, height)

	// Do anything extra for a  view
}
