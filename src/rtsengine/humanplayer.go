package rtsengine

import "image"

// HumanPlayer implements the IPlayer interface for a human player
type HumanPlayer struct {
	View

	// Name of this player
	description string
}

// NewHumanPlayer constructs a HumaPlayer
func NewHumanPlayer(description string, worldLocation image.Point, width int, height int) *HumanPlayer {
	player := HumanPlayer{}

	player.description = description
	player.GenerateView(worldLocation, width, height)

	return &player
}

func (player *HumanPlayer) name() string {
	return player.description
}
