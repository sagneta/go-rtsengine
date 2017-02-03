package main

import (
	"encoding/json"
	"fmt"
	"net"
	"rtsengine"
)

func main() {
	fmt.Print("Reference UI")

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
