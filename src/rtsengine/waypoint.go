package rtsengine

import (
	"fmt"
	"image"
)

// Waypoint maintains scoring for a waypoint in the World Grid.
type Waypoint struct {
	Poolable

	// The parent Waypoint in a path
	Parent *Waypoint

	// Location in the World Grid in world coordinates
	Locus image.Point

	// Scoring.
	// g is the cost it takes to get to this Waypoint
	// h is our guess (heuristic) as to how much it'll cost to reach the goal from that node
	// f = g + h so f is the final cost. The lower the better.
	F, G, H float64

	// Position (also known as a neihborhood
	Position int
}

// Print will dump the contents of this Waypoint
func (s *Waypoint) Print() {
	fmt.Printf("Locus(%d,%d) Position(%d) F(%f) G(%f) H(%f)\n", s.Locus.X, s.Locus.Y, s.Position, s.F, s.G, s.H)
}
