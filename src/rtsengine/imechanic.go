package rtsengine

/*
 Interface
 Encapsulates a specific mechanic. Such as a infantry mechanic, a castle mechanic, a pathing mechanic,
 etcetera.
*/

// IMechanic encapsulates are particular mechanic managed by the game.
type IMechanic interface {
	name() string
}
