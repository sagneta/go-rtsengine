package rtsengine

import "fmt"

// WireCommand enumeration
type WireCommand byte

const (
	// NOOP is the no operation
	NOOP WireCommand = iota + 1

	// MoveUnit moves unit to and fro. The rts-engine produces this command
	MoveUnit

	// PathUnitToLocation will set the destination of a unit to the CurrentX, CurrentY
	// does starting a pathing operation for the unit. Used by both client and rts-engine.
	PathUnitToLocation

	// NewUnitAdded means a new unit has been generated and added to the world.
	NewUnitAdded

	// UnitDestroyed means just that. That means dead and should be removed.
	UnitDestroyed

	// FullRefreshPlayerToUI means a screen refresh inwhich every acre of a player's view
	// is sent over the wire.
	FullRefreshPlayerToUI

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

	// WhoAmI queries the rtsengine to return information on the current user connection.
	WhoAmI
)

// WirePacket is a packet of data that can be JSON marshalled/unmarshalled
// and sent over the wire. Naming of the fields is IMPORTANT so careful.
type WirePacket struct {
	Command WireCommand

	// The tarrain at CurrentX and CurrentY
	LocalTerrain Terrain

	// Unit ID. <= 0 if no ID
	UnitID int

	// Is the unique ID of the player that owns the unit.
	// OwnerPlayerID  <= 0 if no ID and thus the unit is owned by nobody.
	OwnerPlayerID int

	// The Type of Unit if any. <=0 means no unit
	Unit UnitType

	// Used in MoveUnit, NewUnitAdded, NewUnitAdded, CancelMove, UnitStateRefresh, UnitDestroyed, ScrollView
	CurrentRow, CurrentColumn int // View Coordinates

	// Used in ResourceUpdate
	Gold, Wood, Food, Stone int

	// UnitStateRefresh
	Life int

	// For the World
	WorldWidth, WorldHeight, WorldRow, WorldColumn int // World Coordinates

	// For the View. The ViewRow and ViewColumn are in world coordinates
	// FullView uses these. SetView sets the new ViewWidth and ViewHeight
	ViewWidth, ViewHeight, ViewRow, ViewColumn int

	// WhoAmI fills this in and nothing more.
	PlayerName string // name of this player
	PlayerID   int    // unique ID of this player

}

// Clear will reinitialize the structure for reuse.
func (p *WirePacket) Clear() {
	p.Command = NOOP
	p.Unit = 0
	p.LocalTerrain = Grass
	p.CurrentRow = 0
	p.CurrentColumn = 0
	p.Gold = 0
	p.Wood = 0
	p.Food = 0
	p.Stone = 0
	p.Life = 0
	p.WorldWidth = 0
	p.WorldHeight = 0
	p.WorldRow = 0
	p.WorldColumn = 0
	p.ViewWidth = 0
	p.ViewHeight = 0
	p.ViewRow = 0
	p.ViewColumn = 0
	p.UnitID = 0

}

// Print will dump the contents of the packet
func (p *WirePacket) Print() {
	fmt.Printf("ID(%d) Command(%d) CurrentX(%d) CurrentY(%d)  Gold(%d) Wood(%d) Food(%d) Stone(%d) Life(%d) Terrain(%d) UnitType(%d)", p.UnitID, p.Command, p.CurrentRow, p.CurrentColumn, p.Gold, p.Wood, p.Food, p.Stone, p.Life, p.LocalTerrain, p.Unit)
}
