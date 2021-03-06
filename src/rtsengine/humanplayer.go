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

	player.Map = make(map[int]IUnit)

	player.AutoNumber.Initialize()

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
		if player.Wire.ReceiveCheckEOF(&packet) {
			fmt.Println("\n\nEOF was detected. Connection lost.")
			return // stops this coroutine
		}
		//packet.Print()
		err := player.dispatch(&packet)
		if err != nil {
			// Stop the cooroutine
			return
		}

	} // for ever
}

func (player *HumanPlayer) dispatch(packet *WirePacket) error {
	switch packet.Command {
	case WhoAmI:
		packet.PlayerName = player.name()
		packet.PlayerID = player.id()

		if err := player.Wire.SendAll(packet); err == io.EOF {
			fmt.Println("\n\nEOF was detected. Connection lost.")
			return err
		}

	case ScrollView:
		X1 := player.WorldOrigin.X + packet.CurrentRow
		Y1 := player.WorldOrigin.Y + packet.CurrentColumn
		point := image.Rect(X1, Y1, X1+player.Span.Dx(), Y1+player.Span.Dy())

		if point.In(player.OurWorld.Span) {
			player.WorldOrigin.X += packet.CurrentRow
			player.WorldOrigin.Y += packet.CurrentColumn
			player.refreshPlayerToUI(true)
		}

	case PathUnitToLocation:
		if packet.UnitID > 0 {
			//fmt.Printf("Move Unit(%d)\n", packet.UnitID)
			unit := player.Map[packet.UnitID]
			destinationWorld := player.ToWorldPoint(&image.Point{packet.CurrentRow, packet.CurrentColumn})
			if unit != nil && player.In(&destinationWorld) {
				unit.movement().MovementDestination = &destinationWorld
			}
		}

	case MoveUnit:
		if player.In(&image.Point{packet.WorldRow, packet.WorldColumn}) {
			if err := player.Wire.SendAll(packet); err == io.EOF {
				fmt.Println("\n\nEOF was detected. Connection lost.")
				return err
			}
		}

	// Set the View to equal the entire world. Used for testing.
	case FullView:
		if err := player.fullView(); err == io.EOF {
			fmt.Println("\n\nEOF was detected. Connection lost.")
			return err
		}

	case SetView:
		// TODO: This needs to be smarter as it could go off the world.
		player.View.Span = image.Rect(0, 0, packet.ViewWidth, packet.ViewHeight)

		// Return all non empty or non grass acres in the view.
	case PartialRefreshPlayerToUI:
		player.refreshPlayerToUI(true)

	case FullRefreshPlayerToUI:
		player.refreshPlayerToUI(false)

	} //switch

	return nil
}

func (player *HumanPlayer) fullView() error {
	packetArray := make([]WirePacket, 1)

	// Set the player to the world coordinates
	player.View.Span = player.OurWorld.Grid.Span
	player.View.WorldOrigin = player.OurWorld.Grid.WorldOrigin

	packetArray[0].Command = FullView
	packetArray[0].ViewRow = 0
	packetArray[0].ViewColumn = 0
	packetArray[0].ViewHeight = player.OurWorld.Grid.Span.Dx()
	packetArray[0].ViewWidth = player.OurWorld.Grid.Span.Dy()

	if err := player.Wire.Send(packetArray); err == io.EOF {
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
			ourAcre := &player.OurWorld.Matrix[worldPoint.X][worldPoint.Y]

			// Partial results are sent only for occupied or non grassy areas.
			if isPartial && !ourAcre.IsOccupiedOrNotGrass() {
				continue
			}

			packet := WirePacket{}

			if isPartial {
				packet.Command = PartialRefreshPlayerToUI
			} else {
				packet.Command = FullRefreshPlayerToUI
			}

			// Use View Coordinates
			packet.CurrentRow = i
			packet.CurrentColumn = j
			packet.LocalTerrain = ourAcre.terrain

			// if occupied use the unit id else use the acre id
			if ourAcre.Occupied() {
				packet.Unit = ourAcre.unit.unitType()
				packet.UnitID = ourAcre.unit.id()
				packet.Life = ourAcre.unit.life()
				if ourAcre.unit.owner() != nil { // Is there an owner?
					packet.OwnerPlayerID = ourAcre.unit.owner().id()
				}
			} else {
				packet.UnitID = ourAcre.id()
			}

			packet.WorldWidth = player.OurWorld.Grid.Span.Dy()
			packet.WorldHeight = player.OurWorld.Grid.Span.Dx()
			packet.WorldRow = 0
			packet.WorldColumn = 0

			packet.ViewWidth = player.View.Span.Dy()
			packet.ViewHeight = player.View.Span.Dx()
			packet.ViewRow = player.View.WorldOrigin.X
			packet.ViewColumn = player.View.WorldOrigin.Y

			packetArray = append(packetArray, packet)

		}
	}

	_ = player.Wire.Send(packetArray)

}
