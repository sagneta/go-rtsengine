package rtsengine

import "fmt"

// BasePlayer maintains all structures common to all kinds of players.
// Presently there are two players {HumanPlayer, AIPlayer}
type BasePlayer struct {
	// The world map that maintains the terrain and units.
	OurWorld *World

	// Our view (projection) onto that world
	View

	// Map of units for this player
	UnitMap

	// Name of this player
	description string

	// Our master pool for frequently used items
	ItemPool *Pool

	// The automated mechanics of this particular user
	Mechanics []IMechanic

	// Pathing systems
	Pathing *AStarPathing
}

// IPlayer Interface
func (player *BasePlayer) name() string {
	return player.description
}

// DumpUnits demonstrates how to do that precisely.
func (player *BasePlayer) DumpUnits() {
	for k, v := range player.AllUnits() {
		fmt.Printf("Player %s UNITS: key[%s] value[%s]\n", player.name(), k, v.name())
	}
}

// PlayerView returns the view associated with this player
func (player *BasePlayer) PlayerView() *View {
	return &player.View
}

// PlayerUnits returns the player unit map for this player.
func (player *BasePlayer) PlayerUnits() *UnitMap {
	return &player.UnitMap
}
