package rtsengine

import (
	"image"
	"math/rand"
	"time"
)

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

	// Generator Random number generator for this view
	Generator *rand.Rand
}

// GenerateView will initialize all internal structures.
// It will set the grid widith and height and situate the
// view onto the world at worldLocation
func (view *View) GenerateView(worldLocation image.Point, width int, height int) {
	view.WorldOrigin = worldLocation
	view.Span = image.Rect(0, 0, height, width)

	view.Generator = rand.New(rand.NewSource(time.Now().UnixNano()))
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

// CenterOfRect returns the center of the Rectangle passed as a parameter.
func (view *View) CenterOfRect(rect *image.Rectangle) image.Point {
	return image.Point{rect.Min.X + (rect.Dx() / 2), rect.Min.Y + (rect.Dy() / 2)}
}

//RandomPointInView returns a pointer to a point randomly selected within the view.
func (view *View) RandomPointInView() *image.Point {
	return &image.Point{view.Generator.Intn(view.Span.Max.X), view.Generator.Intn(view.Span.Max.Y)}
}

// RandomPointClostToPoint will generate a point close to locus no farther than maxDistance away
func (view *View) RandomPointClostToPoint(locus *image.Point, maxDistance int) *image.Point {
	return &image.Point{locus.X + view.Generator.Intn(maxDistance), locus.Y + view.Generator.Intn(maxDistance)}
}
