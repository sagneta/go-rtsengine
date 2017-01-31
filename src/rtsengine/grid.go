package rtsengine

import (
	"errors"
	"image"
)

// Grid maintains an acre grid and its span.
type Grid struct {
	// Actual data copy of a portion of the world grid
	Matrix [][]Acre

	// Width and Height of this Grid
	Span image.Rectangle

	// Where the upper left hand corner of this grid
	// is located in world coordinates. If it is 0,0 then
	// WorldOrigin == Grid
	WorldOrigin image.Point
}

// GenerateGrid will initialize all internal structures.
// It will set the grid widith and height and situate the
// grid onto the world at worldLocation
func (grid *Grid) GenerateGrid(worldLocation image.Point, width int, height int) {
	grid.WorldOrigin = worldLocation
	grid.Span = image.Rect(0, 0, width, height)

	// allocate 2d array row per row.
	grid.Matrix = make([][]Acre, height)
	for i := range grid.Matrix {
		grid.Matrix[i] = make([]Acre, width)
	}

}

// ToGridPoint Converts world coordinates to grid coordinates
func (grid *Grid) ToGridPoint(worldPoint image.Point) image.Point {
	return worldPoint.Sub(grid.WorldOrigin)
}

// ToWorldPoint converts grid coordinates to world coordinates
func (grid *Grid) ToWorldPoint(gridPoint image.Point) image.Point {
	return gridPoint.Add(grid.WorldOrigin)
}

// In returns true if worldPoint is In the grid. False otherwise.
func (grid *Grid) In(worldPoint image.Point) bool {
	return grid.ToGridPoint(worldPoint).In(grid.Span)
}

// Overlaps returns true if the other grid overlaps with this grid
func (grid *Grid) Overlaps(other *Grid) bool {
	return grid.Span.Overlaps(other.Span)
}

// Remove will eliminate a unit from the grid where-ever it is fine.
// The algorithm presently is brute force.
func (grid *Grid) Remove(unit IUnit) {
	for i := range grid.Matrix {
		for j := range grid.Matrix[i] {
			if grid.Matrix[i][j].unit == unit {
				grid.Matrix[i][j].unit = nil
			}
		}
	}
}

// Add will place the unit in the grid at location. Error is returned
// if the location is invalid. That is outside the known grid.
func (grid *Grid) Add(unit IUnit, location image.Point) error {
	if !grid.In(location) {
		return errors.New("Location not within the world")
	}

	grid.Matrix[location.X][location.Y].unit = unit

	return nil
}
