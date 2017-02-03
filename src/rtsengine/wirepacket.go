package rtsengine

import "fmt"

// WireCommand enumeration
type WireCommand byte

const (
	// MoveUnit moves unit to and fro
	MoveUnit = iota + 1

	// NewUnitAdded means a new unit has been generated and added to the world.
	NewUnitAdded

	// UnitDestroyed means just that. That means dead and should be removed.
	UnitDestroyed

	// RefreshPlayerToUI means a screen refresh inwhich every acre of a player's view
	// is sent over the wire.
	RefreshPlayerToUI

	// CancelMove will cancel a move in progress
	CancelMove

	// UnitStateRefresh means the complete state of a unit is sent over the wire to the ui.
	UnitStateRefresh

	// ResourceUpdate means the current tally of resources of a player should be sent over the wire to the ui.
	ResourceUpdate
)

// WirePacket is a packet of data that can be JSON marshalled/unmarshalled
// and sent over the wire. Naming of the fields is IMPORTANT so careful.
type WirePacket struct {
	Command WireCommand

	// Used in MoveUnit, NewUnitAdded, NewUnitAdded, CancelMove, UnitStateRefresh, UnitDestroyed
	CurrentX, CurrentY int

	// Used in MoveUnit Command
	ToX, ToY int

	// Used in ResourceUpdate
	Gold, Wood, Food, Stone int

	// UnitStateRefresh
	Life int
}

// Print will dump the contents of the packet
func (p *WirePacket) Print() {
	fmt.Printf("Command(%d) CurrentX(%d) CurrentY(%d) ToX(%d) ToY(%d) Gold(%d) Wood(%d) Food(%d) Stone(%d) Life(%d)", p.Command, p.CurrentX, p.CurrentY, p.ToX, p.ToY, p.Gold, p.Wood, p.Food, p.Stone, p.Life)
}
