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

	// Initialized UDPWire to listen upon
	listen(wire *UDPWire)

	// TRUE if human player
	isHuman() bool

	// TRUE if current has network connection
	isWireAlive() bool

	// Invoke to cause the player to begin play immediately.
	start()

	// Invoke to cause the player to stop play immediately.
	stop()
}
