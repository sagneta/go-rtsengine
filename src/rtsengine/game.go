package rtsengine

// Game is an actual game with UDP ports and IPlayers
// In theory the rtsengine can maintain N number of simultaneous
// running Games as long as UDP ports do not overlap.
type Game struct {
	// Our players for this game.
	// Once the game begins this array does not change.
	Players []IPlayer

	// The world map that maintains the terrain and units.
	OurWorld World

	// The automated mechanics of this particular game.
	Mechanics []IMechanic
}
