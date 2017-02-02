package rtsengine

import "container/list"

// GameMaster maintains an array of games.
// The rtsengine can run N number of simultaneous Games each
// with N number of players.
type GameMaster struct {
	GamesList *list.List
}

// NewGameMaster constructs a new game master
func NewGameMaster() *GameMaster {

	gm := GameMaster{}

	gm.GamesList = list.New()

	return &gm
}

// Add will add the game to the Games Master.
func (gameMaster *GameMaster) Add(game *Game) {
	gameMaster.GamesList.PushBack(game)
}
