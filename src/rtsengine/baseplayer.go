package rtsengine

import "fmt"

// BasePlayer maintains all structures common to all kinds of players.
// Presently there are two players {HumanPlayer, AIPlayer}
type BasePlayer struct {
	View
	UnitMap

	// Name of this player
	description string

	// The automated mechanics of this particular user
	Mechanics []IMechanic
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
