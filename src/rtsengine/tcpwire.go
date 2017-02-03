package rtsengine

import (
	"encoding/json"
	"net"
)

/*
 Encapsulates the TCP wire connection between this server and a UI client.
*/

// TCPWire encapsulates our TCP wire connection.
type TCPWire struct {
	Connection  net.Conn
	JSONDecoder *json.Decoder
	JSONEncoder *json.Encoder
}
