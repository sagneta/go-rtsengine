package rtsengine

import "image"

// HumanPlayer implements the IPlayer interface for a human player
type HumanPlayer struct {
	// Structures common to all players.
	BasePlayer

	// Live UDPWire to communicate with UI
	Wire *UDPWire
}

// NewHumanPlayer constructs a HumanPlayer
func NewHumanPlayer(description string, worldLocation image.Point, width int, height int) *HumanPlayer {
	player := HumanPlayer{}

	player.description = description
	player.GenerateView(worldLocation, width, height)

	// Add mechanics

	return &player
}

/////////////////////////////////////////////////////////////////////////
//                           IPlayer interface                         //
/////////////////////////////////////////////////////////////////////////
func (player *HumanPlayer) listen(wire *UDPWire) {
	player.Wire = wire
}

func (player *HumanPlayer) isHuman() bool {
	return true
}

func (player *HumanPlayer) isWireAlive() bool {
	// Best guess
	return player.Wire != nil
}

/////////////////////////////////////////////////////////////////////////
//                           IPlayer interface                         //
/////////////////////////////////////////////////////////////////////////
