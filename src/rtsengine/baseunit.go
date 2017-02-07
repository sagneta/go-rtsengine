package rtsengine

import (
	"fmt"
	"math"
)

// BaseUnit composes all the structures necessary for any unit.
type BaseUnit struct {
	Poolable
	AutoNumber
	HealthAndAttack
	owner IPlayer
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

func (unit *BaseUnit) newWirePacket(command WireCommand, channel chan *WirePacket) *WirePacket {
	packet := WirePacket{}

	switch command {
	case MoveUnit:
		packet.Command = command
		viewPoint := unit.owner.PlayerView().ToViewPoint(unit.CurrentLocation)

		fmt.Print("VIEWPOINT: ")
		fmt.Println(viewPoint)
		packet.CurrentX = viewPoint.X
		packet.CurrentY = viewPoint.Y
		packet.ToX = viewPoint.X
		packet.ToY = viewPoint.Y

		packet.UnitID = unit.id()

		packet.ViewX = unit.owner.PlayerView().WorldOrigin.X
		packet.ViewY = unit.owner.PlayerView().WorldOrigin.Y
		packet.ViewWidth = unit.owner.PlayerView().Span.Dx()
		packet.ViewHeight = unit.owner.PlayerView().Span.Dx()

		channel <- &packet
	}

	return nil
}
