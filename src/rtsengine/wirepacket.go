package rtsengine

import "fmt"

// WireCommand enumeration
type WireCommand byte

const (
	// NOOP is the no operation
	NOOP WireCommand = iota + 1
	// MoveUnit moves unit to and fro
	MoveUnit

	// NewUnitAdded means a new unit has been generated and added to the world.
	NewUnitAdded

	// UnitDestroyed means just that. That means dead and should be removed.
	UnitDestroyed

	// RefreshPlayerToUI means a screen refresh inwhich every acre of a player's view
	// is sent over the wire.
	RefreshPlayerToUI

	// PartialRefreshPlayerToUI means a screen refresh inwhich every unit and non-grass acre
	// is sent over the wire. Tha means the UI should assume missing acres are grass
	// Simple alogorithm for reducing chatter.
	PartialRefreshPlayerToUI

	// CancelMove will cancel a move in progress
	CancelMove

	// UnitStateRefresh means the complete state of a unit is sent over the wire to the ui.
	UnitStateRefresh

	// ResourceUpdate means the current tally of resources of a player should be sent over the wire to the ui.
	ResourceUpdate

	// ScrollView scrolls the view to the new x and y which is ToX and ToY
	ScrollView

	// SetView will set the view directly to world coordinates WorldX, WorldY, Width and Height.
	SetView

	// FullView will set the view to the entire World. Used mostly for testing.
	FullView
)

// WirePacket is a packet of data that can be JSON marshalled/unmarshalled
// and sent over the wire. Naming of the fields is IMPORTANT so careful.
type WirePacket struct {
	Command WireCommand

	// The tarrain at CurrentX and CurrentY
	LocalTerrain Terrain

	// The Type of Unit if any. <=0 means no unit
	Unit UnitType

	// Used in MoveUnit, NewUnitAdded, NewUnitAdded, CancelMove, UnitStateRefresh, UnitDestroyed
	CurrentX, CurrentY int // View Coordinates

	// Used in MoveUnit and ScrollView Command
	ToX, ToY int // View Coordinates

	// Used in ResourceUpdate
	Gold, Wood, Food, Stone int

	// UnitStateRefresh
	Life int

	// For the World
	WorldWidth, WorldHeight, WorldX, WorldY int // World Coordinates

	// For the View. The ViewX and ViewY are in world coordinates
	ViewWidth, ViewHeight, ViewX, ViewY int
}

// Clear will reinitialize the structure for reuse.
func (p *WirePacket) Clear() {
	p.Command = NOOP
	p.Unit = 0
	p.LocalTerrain = Grass
	p.CurrentX = 0
	p.CurrentY = 0
	p.ToX = 0
	p.ToY = 0
	p.Gold = 0
	p.Wood = 0
	p.Food = 0
	p.Stone = 0
	p.Life = 0
	p.WorldWidth = 0
	p.WorldHeight = 0
	p.WorldX = 0
	p.WorldY = 0
	p.ViewWidth = 0
	p.ViewHeight = 0
	p.ViewX = 0
	p.ViewY = 0

}

// Print will dump the contents of the packet
func (p *WirePacket) Print() {
	fmt.Printf("Command(%d) CurrentX(%d) CurrentY(%d) ToX(%d) ToY(%d) Gold(%d) Wood(%d) Food(%d) Stone(%d) Life(%d) Terrain(%d) UnitType(%d)", p.Command, p.CurrentX, p.CurrentY, p.ToX, p.ToY, p.Gold, p.Wood, p.Food, p.Stone, p.Life, p.LocalTerrain, p.Unit)
}
