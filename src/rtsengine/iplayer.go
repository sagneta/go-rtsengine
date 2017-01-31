package rtsengine

/*
 Interface
 A Human or AI Player.
 Will maintain the View and the Unit list for a player.
 All player state resides here.
*/

// IPlayer encapsulates are particular mechanic managed by the game.
type IPlayer interface {
	name() string
}
