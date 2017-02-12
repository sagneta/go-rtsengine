package rtsengine

import "image"

// AIPlayer will implement and Machine Intelligent Player
type AIPlayer struct {
	// Structures common to all players.
	BasePlayer
}

// NewAIPlayer constructs a AIPlayer
func NewAIPlayer(description string, worldLocation image.Point, width int, height int, pool *Pool, pathing *AStarPathing, world *World) *AIPlayer {
	player := AIPlayer{}

	player.description = description
	player.GenerateView(worldLocation, width, height)
	player.ItemPool = pool
	player.Pathing = pathing
	player.OurWorld = world

	player.AutoNumber.Initialize()
	// Add mechanics

	return &player
}

/////////////////////////////////////////////////////////////////////////
//                           IPlayer interface                         //
/////////////////////////////////////////////////////////////////////////
func (player *AIPlayer) listen(wire *TCPWire) {
	// Does nothing beyond satisfying the interface.
}

func (player *AIPlayer) isHuman() bool {
	return false
}

func (player *AIPlayer) isWireAlive() bool {
	return false
}

func (player *AIPlayer) start() error {

	return nil
}

func (player *AIPlayer) stop() {

}

func (player *AIPlayer) dispatch(packet *WirePacket) error {
	return nil
}

/////////////////////////////////////////////////////////////////////////
//                           IPlayer interface                         //
/////////////////////////////////////////////////////////////////////////
