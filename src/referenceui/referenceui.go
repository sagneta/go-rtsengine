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
	termbox "github.com/nsf/termbox-go"

	tl "github.com/JoelOtter/termloop"
)

// ScreenController on our screen
type ScreenController struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel

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

// Assume grass is the default.
func main() {
	testNetwork()

	time.Sleep(time.Second * 60)

	// Ruin network communication etcetera
	castle, _ := utf8.DecodeRuneInString("\xF0\x9F\x8F\xAF")
	//fmt.Printf("Reference UI \xF0\x9F\x8F\xAF  Castle %c", castle)

	// A Canvas is a 2D array of Cells, used for drawing.
	// The structure of a Canvas is an array of columns.
	// This is so it can be addrssed canvas[x][y].
	type Canvas [][]termbox.Cell

	game := termloop.NewGame()

	screen := game.Screen()
	screenWidth, screenHeight := screen.Size()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: '.',
	})

	screenController := ScreenController{
		Entity:       tl.NewEntity(1, 1, 1, 1),
		level:        level,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
	// Set the character at position (0, 0) on the entity.
	screenController.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: castle})
	level.AddEntity(&screenController)

	// Lake
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))

	game.Screen().SetLevel(level)
	game.Start()

}

func testNetwork() {
	var packet rtsengine.WirePacket

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer conn.Close()

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	packet.Command = rtsengine.FullView
	err = encoder.Encode(&packet)
	if err != nil {
		fmt.Println("Unexpected wire error", err)
		return
	}

	var packetArray []rtsengine.WirePacket
	if err := decoder.Decode(&packetArray); err == io.EOF {
		fmt.Println("\n\nEOF was detected. Connection lost.")
		return
	}

	for _, p := range packetArray {
		fmt.Printf("Set View to width(%d) height(%d)\n\n", p.ViewWidth, p.ViewHeight)
	}

	packet.Command = rtsengine.PartialRefreshPlayerToUI
	packet.ToX = -1
	packet.ToY = -2

	err = encoder.Encode(&packet)
	if err != nil {
		fmt.Println("Unexpected wire error", err)
		return
	}

	var packetArray2 []rtsengine.WirePacket
	if err := decoder.Decode(&packetArray2); err == io.EOF {
		fmt.Println("\n\nEOF was detected. Connection lost.")
		return
	}

	for _, p := range packetArray2 {
		p.Print()
		fmt.Println("")
	}

}
