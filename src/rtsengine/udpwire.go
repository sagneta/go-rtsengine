package rtsengine

import "net"

/*
 Encapsulates the UDP wire connection between this server and a UI client.
*/

// UDPWire encapsulates our UDP wire connection.
type UDPWire struct {
	Host       string
	Port       string
	Connection net.PacketConn
}
