package rtsengine

import (
	"encoding/json"
	"io"
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

// Receive wire packet over the wire and returns any error
func (wire *TCPWire) Receive(packet *WirePacket) error {
	return wire.JSONDecoder.Decode(packet)
}

// ReceiveCheckEOF accepts packet over the wire and returns TRUE
// if EOF encountered
func (wire *TCPWire) ReceiveCheckEOF(packet *WirePacket) bool {
	return wire.JSONDecoder.Decode(packet) == io.EOF
}

// Send a packetarray over the wire. Returns any error
func (wire *TCPWire) Send(packetArray []WirePacket) error {
	return wire.JSONEncoder.Encode(&packetArray)
}

// SendCheckEOF will send packetArray over the wire.
// If send failed and EOF detected TRUE is returned.
func (wire *TCPWire) SendCheckEOF(packetArray []WirePacket) bool {
	return wire.JSONEncoder.Encode(&packetArray) == io.EOF
}
