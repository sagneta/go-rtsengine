package rtsengine

// Square maintains scoring for a square in the World Grid.
type Square struct {
	// The parent Square in a path
	parent *Square

	// Location in the World Grid in world coordinates
	x, y int

	// Scoring.
	// g is the cost it takes to get to this Square
	// h is our guess (heuristic) as to how much it'll cost to reach the goal from that node
	// f = g + h so f is the final cost. The lower the better.
	f, g, h int
}
