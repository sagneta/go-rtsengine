package rtsengine

import "math"

// BaseUnit composes all the structures necessary for any unit.
type BaseUnit struct {
	Poolable
	AutoNumber
	HealthAndAttack
	Owner IPlayer
	Movement
}

// Initialize will set the unit to the base state.
// Call only once per instantiation
func (unit *BaseUnit) Initialize() {
	unit.Deallocate()
	unit.AutoNumber.Initialize()

	// Default to something huge which make it immovable.
	// Override if you want the unit to move.
	unit.DeltaInMillis = math.MaxInt64
}

// IUnit
func (unit *BaseUnit) movement() *Movement {
	return &unit.Movement
}

func (unit *BaseUnit) owner() IPlayer {
	return unit.Owner
}

// sendPacketToChannel will do just that. :)
func (unit *BaseUnit) sendPacketToChannel(command WireCommand, channel chan *WirePacket) {
	switch command {
	case MoveUnit:
		channel <- unit.fillOutUnit(&WirePacket{}, command)
	}
}

// fillOutUnit will fill out the packet and return a reference to that same packet.
// This method does most of the work. Some command specific mutations will need to be
// done elsewhere.
func (unit *BaseUnit) fillOutUnit(packet *WirePacket, command WireCommand) *WirePacket {
	packet.Command = command
	viewPoint := unit.Owner.PlayerView().ToViewPoint(unit.CurrentLocation)

	packet.CurrentX = viewPoint.X
	packet.CurrentY = viewPoint.Y

	packet.UnitID = unit.id()

	packet.ViewX = unit.Owner.PlayerView().WorldOrigin.X
	packet.ViewY = unit.Owner.PlayerView().WorldOrigin.Y
	packet.ViewWidth = unit.Owner.PlayerView().Span.Dx()
	packet.ViewHeight = unit.Owner.PlayerView().Span.Dx()

	packet.Life = unit.Life

	packet.WorldX = unit.CurrentLocation.X
	packet.WorldY = unit.CurrentLocation.Y
	return packet
}
