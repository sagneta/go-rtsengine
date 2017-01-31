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

// GenerateView will initialize all internal structures.
// It will set the grid widith and height and situate the
// view onto the world at worldLocation
func (view *View) GenerateView(worldLocation image.Point, width int, height int) {
	view.worldOrigin = worldLocation
	view.span = image.Rect(0, 0, width, height)

	// allocate 2d array row per row.
	view.Grid = make([][]Acre, height)
	for i := range view.Grid {
		view.Grid[i] = make([]Acre, width)
	}

}

// Converts world coordinates to view coordinates
func (view *View) toViewPoint(worldPoint image.Point) image.Point {
	return worldPoint.Sub(view.worldOrigin)
}

// In returns true if worldPoint is In the view. False otherwise.
func (view *View) In(worldPoint image.Point) bool {
	return view.toViewPoint(worldPoint).In(view.span)
}

// Overlaps returns true if the other view overlaps with this view
func (view *View) Overlaps(other *View) bool {
	return view.span.Overlaps(other.span)
}
