package rtsengine

import "image"

/*
 A View into the world grid

*/

// View is a projection onto the World Grid
type View struct {
	// Actual data copy of a portion of the world grid
	Grid [][]Acre

	// Width and Height of this Grid
	span image.Rectangle

	// Where the upper left hand corner of this grid
	// is located in world coordinates
	worldOrigin image.Point
}

// Converts world coordinates to view coordinates
func (view *View) toViewPoint(worldPoint image.Point) image.Point {
	return worldPoint.Sub(view.worldOrigin)
}

// In returns true if worldPoint is in the view. False otherwise.
func (view *View) In(worldPoint image.Point) bool {
	return view.toViewPoint(worldPoint).In(view.span)
}
