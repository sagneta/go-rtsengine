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
	if player.isWireAlive() {
		// Will stop the listenForWireCommands coroutine
		err := player.Wire.Connection.Close()
		if err != nil {
			fmt.Println(err)
		}
		player.Wire = nil
	}
}

/////////////////////////////////////////////////////////////////////////
//                           IPlayer interface                         //
/////////////////////////////////////////////////////////////////////////

// listenForWireCommands will listend for commands from a human player
// and will perform the proper command responses.
func (player *HumanPlayer) listenForWireCommands() {
	var packet WirePacket
	for {

		// Blocking call
		if err := player.Wire.JSONDecoder.Decode(&packet); err == io.EOF {
			fmt.Println("\n\nEOF was detected. Connection lost.")
			return // stops this coroutine
		}
		packet.Print()

		switch packet.Command {

		// Set the View to equial the entire world. Used for testing.
		case FullView:
			player.fullView()

			// Return all non empty or non grass acres in the view.
		case PartialRefreshPlayerToUI:
			player.refreshPlayerToUI(true)

		case FullRefreshPlayerToUI:
			player.refreshPlayerToUI(false)

		} //switch

	} // for ever
}

func (player *HumanPlayer) fullView() error {
	packetArray := make([]WirePacket, 1)

	// Set the player to the world coordinates
	player.View.Span = player.OurWorld.Grid.Span
	player.View.WorldOrigin = player.OurWorld.Grid.WorldOrigin

	packetArray[0].Command = FullView
	packetArray[0].ViewX = 0
	packetArray[0].ViewY = 0
	packetArray[0].ViewHeight = player.OurWorld.Grid.Span.Dx()
	packetArray[0].ViewWidth = player.OurWorld.Grid.Span.Dy()

	if err := player.Wire.JSONEncoder.Encode(&packetArray); err == io.EOF {
		fmt.Println("\n\nEOF was detected. Connection lost.")
		return err
	}

	return nil
}

func (player *HumanPlayer) refreshPlayerToUI(isPartial bool) {
	var packetArray []WirePacket

	for i := 0; i < player.View.Span.Dx(); i++ {
		for j := 0; j < player.View.Span.Dy(); j++ {

			// Convert View point to world and get acre in world.
			worldPoint := player.View.ToWorldPoint(&image.Point{i, j})
			if !player.OurWorld.In(&worldPoint) {
				continue
			}
			ourAcre := player.OurWorld.Matrix[worldPoint.X][worldPoint.Y]

			// Partial results are sent only for occupied or non grassy areas.
			if isPartial && !ourAcre.IsOccupiedOrNotGrass() {
				continue
			}

			packet := WirePacket{}
			//packet.Clear()

			if isPartial {
				packet.Command = PartialRefreshPlayerToUI
			} else {
				packet.Command = FullRefreshPlayerToUI
			}

			// Use View Coordinates
			packet.CurrentX = i
			packet.CurrentY = j
			packet.LocalTerrain = ourAcre.terrain

			// if occupied use the unit id else use the acre id
			if ourAcre.Occupied() {
				packet.Unit = ourAcre.unit.unitType()
				packet.UnitID = ourAcre.unit.id()
				packet.Life = ourAcre.unit.life()
			} else {
				packet.UnitID = ourAcre.id()
			}

			packet.WorldWidth = player.OurWorld.Grid.Span.Dy()
			packet.WorldHeight = player.OurWorld.Grid.Span.Dx()
			packet.WorldX = 0
			packet.WorldY = 0

			packet.ViewWidth = player.View.Span.Dy()
			packet.ViewHeight = player.View.Span.Dx()
			packet.ViewX = player.View.WorldOrigin.X
			packet.ViewY = player.View.WorldOrigin.Y

			packetArray = append(packetArray, packet)

		}
	}

	if err := player.Wire.JSONEncoder.Encode(&packetArray); err == io.EOF {
		fmt.Println("\n\nEOF was detected. Connection lost.")
		return
	}

}
