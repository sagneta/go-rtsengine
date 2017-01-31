package rtsengine

/*
 A View into the world grid

*/

// View is a projection onto the World Grid
type View struct {
	// Actual data copy of a portion of the world grid
	Grid [][]Acre

	// Width and Height of this Grid
	width  int
	height int

	// Where the upper left hand corner of this grid
	// is located in world coordinates
	worldX int
	worldY int
}
