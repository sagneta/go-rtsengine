package rtsengine

import (
	"errors"
	"fmt"
	"image"
	"math"
	"math/rand"
	"time"
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

	// Generator Random number generator for this view
	Generator *rand.Rand
}

// GenerateGrid will initialize all internal structures.
// It will set the grid widith and height and situate the
// grid onto the world at worldLocation
func (grid *Grid) GenerateGrid(worldLocation image.Point, width int, height int) {
	grid.WorldOrigin = worldLocation
	grid.Span = image.Rect(0, 0, height, width)
	grid.Generator = rand.New(rand.NewSource(time.Now().UnixNano()))

	// allocate 2d array row per row.
	grid.Matrix = make([][]Acre, height)
	for i := range grid.Matrix {
		grid.Matrix[i] = make([]Acre, width)
	}

	for i := range grid.Matrix {
		for j := range grid.Matrix[i] {
			grid.Matrix[i][j].Initialize()
		}
	}
}

// ToGridPoint Converts world coordinates to grid coordinates
func (grid *Grid) ToGridPoint(worldPoint *image.Point) image.Point {
	return worldPoint.Sub(grid.WorldOrigin)
}

// ToWorldPoint converts grid coordinates to world coordinates
func (grid *Grid) ToWorldPoint(gridPoint *image.Point) image.Point {
	return gridPoint.Add(grid.WorldOrigin)
}

// In returns true if worldPoint is In the grid. False otherwise.
func (grid *Grid) In(worldPoint *image.Point) bool {
	return grid.ToGridPoint(worldPoint).In(grid.Span)
}

// Overlaps returns true if the other grid overlaps with this grid
func (grid *Grid) Overlaps(other *Grid) bool {
	return grid.Span.Overlaps(other.Span)
}

// Remove will eliminate a unit from the grid where-ever it is found.
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

// RemoveAt will remove the unit at location
func (grid *Grid) RemoveAt(unit IUnit, location *image.Point) {
	grid.Matrix[location.X][location.Y].unit = nil
}

// Add will place the unit in the grid at location. Error is returned
// if the location is invalid. That is outside the known grid.
func (grid *Grid) Add(unit IUnit, location *image.Point) error {
	if !grid.In(location) {
		return errors.New("Location not within the world")
	}

	grid.Matrix[location.X][location.Y].unit = unit

	return nil
}

// Set the unit at locus within this grid.
func (grid *Grid) Set(locus *image.Point, unit IUnit) error {
	return grid.Matrix[locus.X][locus.Y].Set(unit)
}

// Collision returns true if the locus is already occupied
// by any other unit OR the terrain is inaccessible such as
// Mountains and Trees.
func (grid *Grid) Collision(locus *image.Point) bool {
	return grid.Matrix[locus.X][locus.Y].Collision()
}

// Distance between two points using a floating point operation.
func (grid *Grid) Distance(source *image.Point, destination *image.Point) float64 {
	x2 := (destination.X - source.X) * (destination.X - source.X)
	y2 := (destination.Y - source.Y) * (destination.Y - source.Y)
	return math.Sqrt(float64(x2 + y2))
}

// DistanceManhattan is described here: http://www.policyalmanac.org/games/heuristics.htm
func (grid *Grid) DistanceManhattan(source *image.Point, destination *image.Point) float64 {
	return 10.0 * (math.Abs(float64(source.X-destination.X)) + math.Abs(float64(source.Y-destination.Y)))
}

// DistanceDiagonelShortcut is described here: http://www.policyalmanac.org/games/heuristics.htm
func (grid *Grid) DistanceDiagonelShortcut(source *image.Point, destination *image.Point) float64 {
	xDistance := math.Abs(float64(source.X - destination.X))
	yDistance := math.Abs(float64(source.Y - destination.Y))
	if xDistance > yDistance {
		return 14.0*yDistance + 10.0*(xDistance-yDistance)
	}

	return 14.0*xDistance + 10.0*(yDistance-xDistance)
}

// DistanceInteger is the distance algorithm using integer arithmetic. Don't use intermediate variables.
func (grid *Grid) DistanceInteger(source *image.Point, destination *image.Point) int {
	return int(grid.SqrtHDU32(uint32((destination.X-source.X)*(destination.X-source.X) + (destination.Y-source.Y)*(destination.Y-source.Y))))
}

// DirectLineBresenham returns a direct line between to points.
// This is an integer implemenation and works on all quadrants
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// Actual implementation https://rosettacode.org/wiki/Bitmap/Bresenham%27s_line_algorithm#Go
func (grid *Grid) DirectLineBresenham(source *image.Point, destination *image.Point) []image.Point {
	x0 := source.X
	y0 := source.Y

	x1 := destination.X
	y1 := destination.Y
	var waypoints = make([]image.Point, 1)

	// implemented straight from WP pseudocode
	dx := x1 - x0
	if dx < 0 {
		dx = -dx
	}
	dy := y1 - y0
	if dy < 0 {
		dy = -dy
	}
	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	for {
		waypoints = append(waypoints, image.Point{x0, y0})
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}

	return waypoints
}

// Print the world as ascii text.
func (grid *Grid) Print() {
	for i := range grid.Matrix {
		for j := range grid.Matrix[i] {

			switch grid.Matrix[i][j].unit.(type) {
			case *Fence:
				fmt.Printf("X")
				continue
			}

			switch grid.Matrix[i][j].terrain {
			case Trees:
				fmt.Printf("T")
			case Mountains:
				fmt.Printf("M")
			case Grass:
				fmt.Printf(".")
			default:
				fmt.Printf(".")
			}
		} //j
		fmt.Println("")
	} //i
}

// SqrtHDU32 is the integer square root for unsigned 32 bit values.
func (grid *Grid) SqrtHDU32(x uint32) uint32 {
	//Using uint guarantees native word width
	var t, b, r uint
	t = uint(x)
	p := uint(1 << 30)
	for p > t {
		p >>= 2
	}
	for ; p != 0; p >>= 2 {
		b = r | p
		r >>= 1
		if t >= b {
			t -= b
			r |= p
		}
	}
	return uint32(r)
}

// Center returns the x,y center of this Grid.
func (grid *Grid) Center() image.Point {
	return image.Point{grid.Span.Min.X + (grid.Span.Dx() / 2), grid.Span.Min.Y + (grid.Span.Dy() / 2)}
}

//RandomPointInGrid returns a pointer to a point randomly selected within the grid.
func (grid *Grid) RandomPointInGrid() *image.Point {
	return &image.Point{grid.Generator.Intn(grid.Span.Max.X), grid.Generator.Intn(grid.Span.Max.Y)}
}
