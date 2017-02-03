package rtsengine

import (
	"fmt"
	"image"
	"io"
)

// HumanPlayer implements the IPlayer interface for a human player
type HumanPlayer struct {
	// Structures common to all players.
	BasePlayer

	// Live TCPWire to communicate with UI
	Wire *TCPWire
}

// NewHumanPlayer constructs a HumanPlayer
func NewHumanPlayer(description string, worldLocation image.Point, width int, height int, pool *Pool, pathing *AStarPathing, world *World) *HumanPlayer {
	player := HumanPlayer{}

	player.description = description
	player.GenerateView(worldLocation, width, height)
	player.ItemPool = pool
	player.Pathing = pathing
	player.OurWorld = world

	// Add mechanics

	return &player
}

/////////////////////////////////////////////////////////////////////////
//                           IPlayer interface                         //
/////////////////////////////////////////////////////////////////////////
func (player *HumanPlayer) listen(wire *TCPWire) {
	player.Wire = wire
}

func (player *HumanPlayer) isHuman() bool {
	return true
}

func (player *HumanPlayer) isWireAlive() bool {
	// Best guess
	return player.Wire != nil
}

func (player *HumanPlayer) start() error {

	if !player.isWireAlive() {
		return fmt.Errorf("Failed: This player does not have an active wire connection.")
	}

	go player.listenForWireCommands()
	return nil
}

func (player *HumanPlayer) stop() {

}

/////////////////////////////////////////////////////////////////////////
//                           IPlayer interface                         //
/////////////////////////////////////////////////////////////////////////

func (player *HumanPlayer) listenForWireCommands() {
	var packet WirePacket
	for {

		if err := player.Wire.JSONDecoder.Decode(&packet); err == io.EOF {
			fmt.Println("\n\nEOF was detected. Connection lost.")
			return
		}
		packet.Print()

		switch packet.Command {
		case FullView:
			player.View.Span = player.OurWorld.Grid.Span
			player.View.WorldOrigin = player.OurWorld.Grid.WorldOrigin

		case PartialRefreshPlayerToUI:
			for i := 0; i < player.View.Span.Dx(); i++ {
				for j := 0; j < player.View.Span.Dy(); j++ {

					// Convert View point to world and get acre in world.
					worldPoint := player.View.ToWorldPoint(&image.Point{i, j})
					if !player.OurWorld.In(&worldPoint) {
						continue
					}
					ourAcre := player.OurWorld.Matrix[worldPoint.X][worldPoint.Y]

					if ourAcre.IsOccupiedOrNotGrass() {
						packet.Clear()

						// Use View Coordinates
						packet.CurrentX = i
						packet.CurrentY = j
						packet.LocalTerrain = ourAcre.terrain
						if ourAcre.Occupied() {
							packet.Unit = ourAcre.unit.unitType()
						}

						//packet.Life = ourAcre.unit.
						packet.WorldWidth = player.OurWorld.Grid.Span.Dy()
						packet.WorldHeight = player.OurWorld.Grid.Span.Dx()
						packet.WorldX = 0
						packet.WorldY = 0

						packet.ViewWidth = player.View.Span.Dy()
						packet.ViewHeight = player.View.Span.Dx()
						packet.ViewX = player.View.WorldOrigin.X
						packet.ViewY = player.View.WorldOrigin.Y

						if err := player.Wire.JSONEncoder.Encode(&packet); err == io.EOF {
							fmt.Println("\n\nEOF was detected. Connection lost.")
							return
						}

					}
				}
			}

		} //switch

	} // for ever

}
