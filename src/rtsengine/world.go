package rtsengine

import (
	"fmt"
	"image"
	"math/rand"
	"time"
)

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

// GenerateSimple will generate a simple world for basic testing.
// Good for testing pathing etcetera.
func (world *World) GenerateSimple() {

	// Make all the world grass!
	for i := range world.Matrix {
		for j := range world.Matrix[i] {
			world.Matrix[i][j].unit = nil
			world.Matrix[i][j].terrain = Grass
		}
	}

	// Randomly dot with trees and mountains
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Trees
	for i := 0; i < 100; i++ {
		xr := r1.Intn(world.Span.Dx())
		yr := r1.Intn(world.Span.Dy())
		world.Matrix[xr][yr].terrain = Trees
	}

	// Mountains
	for i := 0; i < 100; i++ {
		xr := r1.Intn(world.Span.Dx())
		yr := r1.Intn(world.Span.Dy())
		world.Matrix[xr][yr].terrain = Mountains
	}

}

// Print the world as ascii text.
func (world *World) Print() {
	for i := range world.Matrix {
		for j := range world.Matrix[i] {
			switch world.Matrix[i][j].terrain {
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
