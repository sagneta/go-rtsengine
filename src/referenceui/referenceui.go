package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"rtsengine"
	"time"

	"github.com/JoelOtter/termloop"

	tl "github.com/JoelOtter/termloop"
)

// ScreenController on our screen
type ScreenController struct {
	*tl.Entity
	level *tl.BaseLevel

	Xdiff int
	Ydiff int

	screenWidth  int
	screenHeight int

	ui *ReferenceUI

	// Last Unit Selection Information
	UnitID int
	MouseX int
	MouseY int
}

// Tick satisfies Entity interface.
func (controller *ScreenController) Tick(event tl.Event) {
	switch event.Key { // If so, switch on the pressed key.
	case tl.KeyArrowRight:
		controller.Ydiff = 1
		controller.Xdiff = 0
		controller.ui.scrollView()

	case tl.KeyArrowLeft:
		controller.Ydiff = -1
		controller.Xdiff = 0
		controller.ui.scrollView()

	case tl.KeyArrowUp:
		controller.Xdiff = -1
		controller.Ydiff = 0
		controller.ui.scrollView()

	case tl.KeyArrowDown:
		controller.Xdiff = 1
		controller.Ydiff = 0
		controller.ui.scrollView()

	case tl.MouseRight, tl.MouseLeft:
		controller.HandleMouseDown(&event)
	}
	/*
			switch event.Type {
			case tl.EventResize:
				log.Print("resized")
				//fmt.Printf("We resized to width(%d) height(%d)", event.Width, event.Height)
				//panic(nil)
			default:
				break
			}

		}
	*/
}

// HandleMouseDown for entire screen.
func (controller *ScreenController) HandleMouseDown(event *tl.Event) {
	acre := controller.ui.findAcre(event.MouseX, event.MouseY)

	// Does our acre exist?
	if acre != nil {
		if acre.Unit > 0 { // has a Unit?
			// Moveable unit?
			switch acre.Unit {
			case rtsengine.UnitInfantry,
				rtsengine.UnitCavalry,
				rtsengine.UnitPeasant,
				rtsengine.UnitShip,
				rtsengine.UnitCatapult,
				rtsengine.UnitArcher:
				// New Unit Selection?
				//fmt.Println("Found the unit")
				if controller.UnitID != acre.UnitID {
					controller.UnitID = acre.UnitID
					controller.MouseX = event.MouseX
					controller.MouseY = event.MouseY
				}
			}
		} else if controller.UnitID > 0 { // was there a previous selection of a unit?
			controller.pathUnit(event)
		}
	} else if controller.UnitID > 0 { // was there a previous selection of a unit?
		controller.pathUnit(event)
	}
}

// pathUnit will path a unit to the mouse location in event.
func (controller *ScreenController) pathUnit(event *tl.Event) {
	id := controller.UnitID
	controller.UnitID = 0
	controller.MouseX = event.MouseX
	controller.MouseY = event.MouseY
	controller.ui.pathUnitToLocation(id, controller.MouseX, controller.MouseY)
}

// Draw the screen.
func (controller *ScreenController) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()

	controller.screenWidth = screenWidth
	controller.screenHeight = screenHeight
	// We need to make sure and call Draw on the underlying Entity.
	controller.Entity.Draw(screen)
}

// Acre is an entity on our screen.
type Acre struct {
	*tl.Entity
	level *tl.BaseLevel

	// Terrain
	LocalTerrain rtsengine.Terrain

	// Unit information if any
	UnitID int
	Unit   rtsengine.UnitType
	Life   int

	// Screen coordinates
	Column int
	Row    int
}

// ReferenceUI which is the master struct for our Game UI
type ReferenceUI struct {
	// TermLoop
	game             *tl.Game
	level            *tl.BaseLevel
	screenController ScreenController

	// Communication
	Connection  net.Conn
	JSONDecoder *json.Decoder
	JSONEncoder *json.Encoder

	// Current viewable acres
	acreMap map[int]*Acre

	// This player information
	playerName string // name of this player
	playerID   int    // unique ID of this player
}

// Start will instantiate the UI and attempt to talk to a rtsengine.
func (ui *ReferenceUI) Start() {
	//time.Sleep(time.Second * 60)

	ui.acreMap = make(map[int]*Acre)

	ui.game = termloop.NewGame()

	screen := ui.game.Screen()
	screenWidth, screenHeight := screen.Size()
	fmt.Printf("w %d h %d\n", screenWidth, screenHeight)

	ui.level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: ' ',
	})

	ui.screenController = ScreenController{
		Entity:       tl.NewEntity(1, 1, 1, 1),
		level:        ui.level,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		ui:           ui,
	}

	// Set the character at position (0, 0) on the entity.
	ui.screenController.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: ' '})
	ui.level.AddEntity(&ui.screenController)

	// Dial up the rtsengine.
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Print(err)
		return
	}

	ui.Connection = conn
	ui.JSONEncoder = json.NewEncoder(ui.Connection)
	ui.JSONDecoder = json.NewDecoder(ui.Connection)

	// Lake
	//level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))

	// Start the game.
	ui.game.Screen().SetLevel(ui.level)

	go ui.communicationPreamble()
	go ui.listenForWireCommands()

	ui.game.Start()

	//time.Sleep(time.Second * 60)
}

// listenForWireCommands will listen for commands from the rtsengine.
func (ui *ReferenceUI) listenForWireCommands() {
	time.Sleep(time.Second * 2)
	for {
		var packetArray []rtsengine.WirePacket
		if err := ui.JSONDecoder.Decode(&packetArray); err == io.EOF {
			fmt.Println("\n\nEOF was detected. Connection lost.")
			return
		}

		if len(packetArray) == 0 {
			continue
		}

		switch packetArray[0].Command {
		case rtsengine.WhoAmI:
			ui.playerName = packetArray[0].PlayerName
			ui.playerID = packetArray[0].PlayerID

		case rtsengine.MoveUnit:
			acre, ok := ui.acreMap[packetArray[0].UnitID]
			if ok {
				acre.Column = packetArray[0].CurrentColumn
				acre.Row = packetArray[0].CurrentRow
				acre.SetPosition(acre.Column, acre.Row)
			}

		// Set the View to equial the entire world. Used for testing.
		case rtsengine.FullView:
			ui.screenController.screenWidth = packetArray[0].ViewWidth
			ui.screenController.screenHeight = packetArray[0].ViewHeight

		case rtsengine.PartialRefreshPlayerToUI:
			ui.handleRefreshPlayerToUI(packetArray)
		}

	}
}

func (ui *ReferenceUI) handleRefreshPlayerToUI(packetArray []rtsengine.WirePacket) {
	var m map[int]*Acre
	m = make(map[int]*Acre)

	for _, p := range packetArray {
		acre, ok := ui.acreMap[p.UnitID]
		if ok {
			delete(ui.acreMap, p.UnitID)
			acre.Column = p.CurrentColumn
			acre.Row = p.CurrentRow
			acre.SetPosition(acre.Column, acre.Row)
		} else {
			acre = &Acre{
				Entity:       tl.NewEntity(p.CurrentColumn, p.CurrentRow, 1, 1),
				level:        ui.level,
				LocalTerrain: p.LocalTerrain,
				Column:       p.CurrentColumn,
				Row:          p.CurrentRow,
				Life:         p.Life,
				UnitID:       p.UnitID,
				Unit:         p.Unit,
			}

			var cell tl.Cell

			switch acre.LocalTerrain {
			case rtsengine.Snow:
				cell = tl.Cell{Fg: tl.ColorWhite, Ch: ' '}

			case rtsengine.Sand:
				cell = tl.Cell{Fg: tl.ColorYellow, Ch: ' '}

			case rtsengine.Water:
				cell = tl.Cell{Fg: tl.ColorBlue, Ch: 'W'}

			case rtsengine.Mountains:
				//r, _ := utf8.DecodeRuneInString("\xF0\x9F\x97\xBB")
				cell = tl.Cell{Fg: tl.ColorBlack, Ch: '^'}

			case rtsengine.Trees:
				//r, _ := utf8.DecodeRuneInString("\xF0\x9F\x8C\xB2")
				//r, _ := utf8.DecodeRuneInString("\xC2\xAE")
				cell = tl.Cell{Fg: tl.ColorWhite, Ch: 'T'}

			case rtsengine.Dirt:
				cell = tl.Cell{Fg: tl.ColorBlack, Ch: 'D'}
				//fmt.Printf("Acre UnitType: %d", acre.Unit)

			}

			switch acre.Unit {
			case rtsengine.UnitWall:
				cell = tl.Cell{Fg: tl.ColorBlack, Ch: ' '}

			case rtsengine.UnitFarm:
				cell = tl.Cell{Fg: tl.ColorYellow, Ch: 'F'}

			case rtsengine.UnitCavalry:
				cell = tl.Cell{Fg: tl.ColorCyan, Ch: 'Z'}

			case rtsengine.UnitInfantry:
				cell = tl.Cell{Fg: tl.ColorBlue, Ch: 'S'}

			case rtsengine.UnitArmory:
				cell = tl.Cell{Fg: tl.ColorBlue, Ch: 'A'}

			case rtsengine.UnitCastle:
				cell = tl.Cell{Fg: tl.ColorBlue, Ch: 'C'}

			case rtsengine.UnitGoldMine:
				cell = tl.Cell{Fg: tl.ColorYellow, Ch: 'G'}

			case rtsengine.UnitHomeStead:
				cell = tl.Cell{Fg: tl.ColorBlue, Ch: 'H'}

			case rtsengine.UnitPeasant:
				cell = tl.Cell{Fg: tl.ColorBlue, Ch: 'P'}

			case rtsengine.UnitStoneQuarry:
				cell = tl.Cell{Fg: tl.ColorWhite, Ch: 's'}

			case rtsengine.UnitShip:
				cell = tl.Cell{Fg: tl.ColorBlue, Ch: 's'}

			case rtsengine.UnitTower:
				cell = tl.Cell{Fg: tl.ColorBlue, Ch: 't'}

			case rtsengine.UnitWoodPile:
				cell = tl.Cell{Fg: tl.ColorBlack, Ch: 'w'}

			case rtsengine.UnitFence:
				//r, _ := utf8.DecodeRuneInString("\xE2\xAC\x9B")
				cell = tl.Cell{Fg: tl.ColorBlack, Ch: 'X'}
			}
			acre.SetCell(0, 0, &cell)
			ui.level.AddEntity(acre)
		}
		m[p.UnitID] = acre
	}

	// Remove anything left over. That means it is no longer within the view.
	for _, v := range ui.acreMap {
		ui.level.RemoveEntity(v)
	}

	ui.acreMap = m
}

// communicationPreamble will start the communication with the rtsengine.
func (ui *ReferenceUI) communicationPreamble() {
	// give the UI some time to initialize.
	time.Sleep(time.Second * 2)

	var packet rtsengine.WirePacket

	// WhoAmI? Get user information for this connection.
	packet.Command = rtsengine.WhoAmI
	err := ui.JSONEncoder.Encode(&packet)
	if err != nil {
		fmt.Println("Unexpected wire error", err)
		return
	}

	// Send full View to set our UI to the entire view of the game for testing.
	packet.Command = rtsengine.SetView

	// They are flipped for our engine.
	packet.ViewHeight, packet.ViewWidth = ui.game.Screen().Size()
	err = ui.JSONEncoder.Encode(&packet)
	if err != nil {
		fmt.Println("Unexpected wire error", err)
		return
	}

	// Force a partial refresh to place the initial acres and units on our screen.
	packet.Command = rtsengine.PartialRefreshPlayerToUI
	err = ui.JSONEncoder.Encode(&packet)
	if err != nil {
		fmt.Println("Unexpected wire error", err)
		return
	}

}

// pathUnitToLocation will request that unitID be pathed to X,Y
func (ui *ReferenceUI) pathUnitToLocation(UnitID int, X int, Y int) {
	var packet rtsengine.WirePacket

	// Send full View to set our UI to the entire view of the game for testing.
	packet.Command = rtsengine.PathUnitToLocation
	packet.CurrentRow = Y
	packet.CurrentColumn = X
	packet.UnitID = UnitID
	err := ui.JSONEncoder.Encode(&packet)
	if err != nil {
		fmt.Println("Unexpected wire error", err)
		return
	}
}

// scrollView will scroll this player view within the world.
func (ui *ReferenceUI) scrollView() {
	var packet rtsengine.WirePacket

	// Send full View to set our UI to the entire view of the game for testing.
	packet.Command = rtsengine.ScrollView
	packet.CurrentRow = ui.screenController.Xdiff
	packet.CurrentColumn = ui.screenController.Ydiff
	err := ui.JSONEncoder.Encode(&packet)
	if err != nil {
		fmt.Println("Unexpected wire error", err)
		return
	}
}

// findAcre will find the acre at X and Y and return nil if not found.
// Presently this is an O(N) search.
func (ui *ReferenceUI) findAcre(X int, Y int) *Acre {
	for _, v := range ui.acreMap {
		if v.Column == X && v.Row == Y {
			return v
		}
	}
	return nil
}

// Assume grass is the default.
func main() {
	ui := ReferenceUI{}

	ui.Start()
}
