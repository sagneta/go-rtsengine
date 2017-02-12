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

	// unique id for this player
	id() int

	// Initialized TCPWire to listen upon
	listen(wire *TCPWire)

	// TRUE if human player
	isHuman() bool

	// TRUE if current has network connection
	isWireAlive() bool

	// Invoke to cause the player to begin play immediately.
	start() error

	// Invoke to cause the player to stop play immediately.
	stop()

	// Returns this player's view
	PlayerView() *View

	// Returns this player's view
	PlayerUnits() *UnitMap

	// Dispatch a wirepacket to this player
	dispatch(packet *WirePacket) error
}
