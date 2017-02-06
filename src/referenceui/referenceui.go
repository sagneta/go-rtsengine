package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"rtsengine"
	"time"
	"unicode/utf8"

	"github.com/JoelOtter/termloop"

	tl "github.com/JoelOtter/termloop"
)

// ScreenController on our screen
type ScreenController struct {
	*tl.Entity
	level *tl.BaseLevel

	prevX int
	prevY int

	screenWidth  int
	screenHeight int
}

// Tick satisfies Entity interface.
func (player *ScreenController) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		x, y := player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.SetPosition(x+1, y)
		case tl.KeyArrowLeft:
			player.SetPosition(x-1, y)
		case tl.KeyArrowUp:
			player.SetPosition(x, y-1)
		case tl.KeyArrowDown:
			player.SetPosition(x, y+1)
		}
	} else {
		switch event.Type {
		case tl.EventResize:
			log.Print("resized")
			//fmt.Printf("We resized to width(%d) height(%d)", event.Width, event.Height)
			//panic(nil)
		default:
			break
		}

	}
}

// Draw the screen. Allows for scrolling
func (player *ScreenController) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()

	player.screenWidth = screenWidth
	player.screenHeight = screenHeight

	/*
		if player.screenHeight == -1 {
			player.screenWidth = screenWidth
			player.screenHeight = screenHeight
		} else if player.screenHeight != screenHeight {
			panic(nil)
		}

	*/
	//x, y := player.Position()
	//player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	// We need to make sure and call Draw on the underlying Entity.
	player.Entity.Draw(screen)

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
	X int
	Y int
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
}

// Start will instantiate the UI and attempt to talk to a rtsengine.
func (ui *ReferenceUI) Start() {
	//time.Sleep(time.Second * 60)

	ui.acreMap = make(map[int]*Acre)

	ui.game = termloop.NewGame()

	screen := ui.game.Screen()
	screenWidth, screenHeight := screen.Size()

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
	}

	// Set the character at position (0, 0) on the entity.
	ui.screenController.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '.'})
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

		switch packetArray[0].Command {
		case rtsengine.MoveUnit:
			acre, ok := ui.acreMap[packetArray[0].UnitID]
			if ok {
				acre.X = packetArray[0].ToY
				acre.Y = packetArray[0].ToX
				acre.SetPosition(acre.X, acre.Y)
			}

		// Set the View to equial the entire world. Used for testing.
		case rtsengine.FullView:
			ui.screenController.screenWidth = packetArray[0].ViewWidth
			ui.screenController.screenHeight = packetArray[0].ViewHeight

		case rtsengine.PartialRefreshPlayerToUI:
			ui.handleRefreshPlayerToUI(packetArray)
			/*
				for _, p := range packetArray {
					p.Print()
					fmt.Println("")
				}
			*/
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
			acre.X = p.CurrentY
			acre.Y = p.CurrentX
			acre.SetPosition(acre.X, acre.Y)
		} else {
			acre = &Acre{
				Entity:       tl.NewEntity(p.CurrentY, p.CurrentX, 1, 1),
				level:        ui.level,
				LocalTerrain: p.LocalTerrain,
				X:            p.CurrentY,
				Y:            p.CurrentX,
				Life:         p.Life,
				UnitID:       p.UnitID,
				Unit:         p.Unit,
			}

			var cell tl.Cell
			switch acre.LocalTerrain {
			case rtsengine.Mountains:
				//r, _ := utf8.DecodeRuneInString("\xF0\x9F\x97\xBB")
				cell = tl.Cell{Fg: tl.ColorBlack, Ch: '^'}

			case rtsengine.Trees:
				//r, _ := utf8.DecodeRuneInString("\xF0\x9F\x8C\xB2")
				cell = tl.Cell{Fg: tl.ColorWhite, Ch: 'T'}
			}

			switch acre.Unit {
			case rtsengine.UnitInfantry:
				cell = tl.Cell{Fg: tl.ColorBlue, Ch: 'S'}
			case rtsengine.UnitFence:
				r, _ := utf8.DecodeRuneInString("\xE2\xAC\x9B")
				cell = tl.Cell{Fg: tl.ColorBlack, Ch: r}

				//			default:
				//				r, _ := utf8.DecodeRuneInString("\xF0\x9F\x92\x82")
				//				cell = tl.Cell{Fg: tl.ColorRed, Ch: r}

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
	//time.Sleep(time.Second * 2)

	var packet rtsengine.WirePacket

	// Send full View to set our UI to the entire view of the game for testing.
	packet.Command = rtsengine.FullView
	err := ui.JSONEncoder.Encode(&packet)
	if err != nil {
		fmt.Println("Unexpected wire error", err)
		return
	}

	// Force a partial refresh to place the initial acres and units on our screen.
	packet.Command = rtsengine.PartialRefreshPlayerToUI
	packet.ToX = -1
	packet.ToY = -2
	err = ui.JSONEncoder.Encode(&packet)
	if err != nil {
		fmt.Println("Unexpected wire error", err)
		return
	}

}

// Assume grass is the default.
func main() {
	ui := ReferenceUI{}

	ui.Start()
}
