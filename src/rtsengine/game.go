package rtsengine

import (
	"fmt"
	"image"
)

// Game is an actual game with UDP ports and IPlayers
// In theory the rtsengine can maintain N number of simultaneous
// running Games as long as UDP ports do not overlap.
type Game struct {
	// Description of game
	Description string

	// Our players for this game.
	// Once the game begins this array does not change.
	Players []IPlayer

	// The world map that maintains the terrain and units.
	OurWorld *World

	// The automated mechanics of this particular game.
	Mechanics []IMechanic

	// Our master pool for frequently used items
	ItemPool *Pool

	// Pathing systems
	Pathing *AStarPathing
}

// NewGame constructs a new game according to the parameters.
func NewGame(
	description string,

	// How many items to pool for decreased GC
	poolItems int,

	noOfHumanPlayers int,
	noOfAIPlayers int,

	// Width and Height of Player View
	playerViewWidth int, playerViewHeight int,

	// Width and Height in Acres of our world.
	worldWidth int, worldHeight int) *Game {

	// This instance
	game := Game{}

	// Item Pool
	game.ItemPool = &Pool{}
	game.ItemPool.Generate(poolItems)

	// Instantiate the pathing system
	game.Pathing = &AStarPathing{}

	// Used for display so we have some idea what games are being played.
	// Make this very descriptive and long. Like '4 Human Players, Fog of War, World(500,500)'
	game.Description = description

	// Instantiate the world
	game.OurWorld = NewWorld(worldWidth, worldHeight)

	// Generate a world. Fill it with trees and rivers and ...
	game.OurWorld.GenerateSimple()

	// Create Players
	game.Players = make([]IPlayer, noOfAIPlayers+noOfHumanPlayers)
	// Create Human Players
	i := 0
	for ; i < noOfHumanPlayers; i++ {
		// The world point needs to be inserted into a random location
		game.Players[i] = NewHumanPlayer(fmt.Sprintf("Human Player %d", i), image.Point{0, 0}, playerViewWidth, playerViewHeight, game.ItemPool, game.Pathing)
	}

	// Create Machine Intelligent Players
	for j := 0; j < noOfAIPlayers; j++ {
		// The world point needs to be inserted into a random location
		game.Players[i] = NewAIPlayer(fmt.Sprintf("AI Player %d", j), image.Point{0, 0}, playerViewWidth, playerViewHeight, game.ItemPool, game.Pathing)
	}

	// Add mechanics

	return &game
}

// Start will start the game.
func (game *Game) Start() {
	for _, player := range game.Players {
		player.start()
	}
}

// Stop will stop the game.
func (game *Game) Stop() {
	for _, player := range game.Players {
		player.stop()
	}
}

// ReadyToGo returns true if we are ready to start a game.
func (game *Game) ReadyToGo() bool {

	// Essentially check if all human players are ready to go.
	// AI's are always ready.
	for _, player := range game.Players {
		if player.isHuman() && !player.isWireAlive() {
			return false
		}
	}

	return true
}
