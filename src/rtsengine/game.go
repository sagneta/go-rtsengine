package rtsengine

import (
	"fmt"
	"image"
	"math/rand"
	"time"
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
	worldWidth int, worldHeight int) (*Game, error) {

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

	// Situate player bases onto the world without overlapping.
	rects, error := game.SituateHomeBases(noOfAIPlayers+noOfHumanPlayers, playerViewWidth, playerViewHeight)

	if error != nil {
		return nil, fmt.Errorf("Failed to situate home bases into world grid. Please reduce number of players and/or increase world size")
	}

	// Create Human Players
	i := 0
	for ; i < noOfHumanPlayers; i++ {
		// The world point needs to be inserted into a random location
		game.Players[i] = NewHumanPlayer(fmt.Sprintf("Human Player %d", i), rects[i].Min, playerViewWidth, playerViewHeight, game.ItemPool, game.Pathing)
	}

	// Create Machine Intelligent Players
	for j := 0; j < noOfAIPlayers; j++ {
		// The world point needs to be inserted into a random location
		game.Players[i] = NewAIPlayer(fmt.Sprintf("AI Player %d", j), rects[i].Min, playerViewWidth, playerViewHeight, game.ItemPool, game.Pathing)
		i++
	}

	// Add mechanics

	return &game, nil
}

// SituateHomeBases will construct home bases in the proper
// locations on the world. That is within the world but not overlapping one another.
// It's possible for large numbers of players on a too small grid this heuristic will not converge
// and an error will be returned.
func (game *Game) SituateHomeBases(noOfPlayers int, playerViewWidth int, playerViewHeight int) ([]*image.Rectangle, error) {
	playerRects := make([]*image.Rectangle, noOfPlayers)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

OUTER:
	for i, j := 0, 0; i < noOfPlayers; j++ {

		// No convergence?
		if j >= 1000 {
			return nil, fmt.Errorf("Not enough space in world grid to insert player grids.")
		}

		// Random point within the world
		randomRect := image.Rect(r1.Intn(game.OurWorld.Span.Dx()), r1.Intn(game.OurWorld.Span.Dy()), playerViewHeight, playerViewWidth)

		// If no players yet just add it and continue.
		if i == 0 {
			playerRects[i] = &randomRect
			i++
			continue
		}

		// Ensure no overlaps with existing player rects
		for _, r := range playerRects {
			// End of array.
			if r == nil {
				break
			}

			// two player home grids overlap. Try again...
			if r.Overlaps(randomRect) {
				continue OUTER
			}
		}

		// no overlap!
		playerRects[i] = &randomRect
		i++
	}

	return playerRects, nil
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