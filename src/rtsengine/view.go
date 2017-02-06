package rtsengine

import "image"

/*
 A View into the world grid

*/

// View is a projection onto the World Grid
type View struct {
	// Width and Height of this Grid
	Span image.Rectangle

	// Where the upper left hand corner of this grid
	// is located in world coordinates. If it is 0,0 then
	// WorldOrigin == Grid
	WorldOrigin image.Point
}

// GenerateView will initialize all internal structures.
// It will set the grid widith and height and situate the
// view onto the world at worldLocation
func (view *View) GenerateView(worldLocation image.Point, width int, height int) {
	view.WorldOrigin = worldLocation
	view.Span = image.Rect(0, 0, height, width)
}

// ToViewPoint Converts world coordinates to view coordinates
func (view *View) ToViewPoint(worldPoint *image.Point) image.Point {
	return worldPoint.Sub(view.WorldOrigin)
}

// ToWorldPoint converts view coordinates to world coordinates
func (view *View) ToWorldPoint(viewPoint *image.Point) image.Point {
	return viewPoint.Add(view.WorldOrigin)
}

// In returns true if worldPoint is In the view. False otherwise.
func (view *View) In(worldPoint *image.Point) bool {
	return view.ToViewPoint(worldPoint).In(view.Span)
}

// Overlaps returns true if the other view overlaps with this view
func (view *View) Overlaps(other *View) bool {
	return view.Span.Overlaps(other.Span)
}

// Center returns the x,y center of this View.
func (view *View) Center() image.Point {
	return image.Point{view.Span.Min.X + (view.Span.Dx() / 2), view.Span.Min.Y + (view.Span.Dy() / 2)}
}
