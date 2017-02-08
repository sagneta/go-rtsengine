package rtsengine

/*
 Interface
 A Unit witin the game which is broader than you might imagine.
 {Gold, Stone, Fence, Wall, Tower, Farm, Castle, Ship, Cavalry, Wood, Catapult, Archer, etcetera}
*/

// UnitType is the type of unit
type UnitType byte

const (
	// UnitArcher or other ranged unit
	UnitArcher UnitType = iota + 1

	// UnitArmory produces soldiers
	UnitArmory
	// UnitCastle fortification
	UnitCastle
	// UnitCatapult or other artillery
	UnitCatapult
	// UnitCavalry or mobile units
	UnitCavalry
	// UnitFarm or food production
	UnitFarm
	// UnitFence wood pallasade or abitus
	UnitFence
	// UnitGoldMine for gold
	UnitGoldMine
	// UnitHomeStead for creating a population
	UnitHomeStead
	// UnitInfantry or foot soldiers
	UnitInfantry
	// UnitPeasant is any villager
	UnitPeasant
	// UnitShip or some military vessel
	UnitShip
	// UnitStoneQuarry produces building material.
	UnitStoneQuarry
	// UnitTower or scout position
	UnitTower
	// UnitWall or defensive parameter
	UnitWall
	// UnitWoodPile or wood production
	UnitWoodPile
)

// IUnit is an interface to all unit within the World and can reside within an Acre.
type IUnit interface {
	IPoolable
	name() string
	unitType() UnitType
	id() int
	life() int
	movement() *Movement
	sendPacketToChannel(command WireCommand, channel chan *WirePacket)
}
