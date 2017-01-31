package rtsengine

/*
 Interface
 A Unit witin the game which is broader than you might imagine.
 {Gold, Stone, Fence, Wall, Tower, Farm, Castle, Ship, Cavalry, Wood, Catapult, Archer, etcetera}
*/

// IUnit is an interface to all unit within the World and can reside within an Acre.
type IUnit interface {
	IPoolable
	name() string
}
