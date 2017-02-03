package main

import (
	"encoding/json"
	"fmt"
	"net"
	"rtsengine"
	"unicode/utf8"

	termbox "github.com/nsf/termbox-go"
)

func main() {

	uni, _ := utf8.DecodeRuneInString("\xF0\x9F\x8F\xAF")
	fmt.Println(uni)

	fmt.Printf("Reference UI \xF0\x9F\x8F\xAF  Castle %c", uni)

	// A Canvas is a 2D array of Cells, used for drawing.
	// The structure of a Canvas is an array of columns.
	// This is so it can be addrssed canvas[x][y].
	type Canvas [][]termbox.Cell

	var packet rtsengine.WirePacket

	packet.Command = rtsengine.RefreshPlayerToUI
	packet.ToX = -1
	packet.ToY = -2

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer conn.Close()

	encoder := json.NewEncoder(conn)

	err = encoder.Encode(&packet)
	if err != nil {
		fmt.Println("Unexpected wire error", err)
		return
	}

}
