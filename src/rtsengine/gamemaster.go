package rtsengine

// GameMaster maintains an array of games.
// The rtsengine can run N number of simultaneous Games each
// with N number of players
type GameMaster struct {
	Games []Game
}
