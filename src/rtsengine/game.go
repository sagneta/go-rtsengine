package rtsengine

import (
	"container/list"
	"encoding/json"
	"fmt"
	"image"
	"net"
	"os"

	"github.com/salviati/go-tmx/tmx"
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

	// Command channel
	CommandChannel chan *WirePacket

	// Our TMX map that describes the world.
	// We never load the images of course.
	TMXMap *tmx.Map

	// First and Last Global Tile Identifier for mountains
	MountainsFirstGID int
	MountainsLastGID  int

	// First and Last Global Tile Identifier for grass
	GrassFirstGID int
	GrassLastGID  int

	// First and Last Global Tile Identifier for trees
	TreesFirstGID int
	TreesLastGID  int

	// First and Last Global Tile Identifier for water
	WaterFirstGID int
	WaterLastGID  int

	// First and Last Global Tile Identifier for dirt
	DirtFirstGID int
	DirtLastGID  int

	// First and Last Global Tile Identifier for sand
	SandFirstGID int
	SandLastGID  int

	// First and Last Global Tile Identifier for snow
	SnowFirstGID int
	SnowLastGID  int

	// Spawn locations. Suggested locations for Home bases.
	// They are called spawns or spawn-points in game lingo.
	SpawnLocations []tmx.Object
}

// NewGame constructs a new game according to the parameters.
func NewGame(
	description string,

	// Name/path to a TMX file to load the world.
	// If nil a crappy default world is produced.
	filenameTMX string,

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

	// The command channel that accepts WirePacket commands
	// and performs the necessary operation.
	game.CommandChannel = make(chan *WirePacket, 500)

	// Used for display so we have some idea what games are being played.
	// Make this very descriptive and long. Like '4 Human Players, Fog of War, World(500,500)'
	game.Description = description

	// Generate a world. Fill it with trees and rivers and ...
	if len(filenameTMX) > 0 {
		tmxmap, err := game.LoadTMX(filenameTMX)
		if err != nil {
			return nil, fmt.Errorf("Could not load TMX file %s", filenameTMX)
		}
		game.TMXMap = tmxmap
		game.RenderTMX()
	} else {
		// Instantiate the world
		game.OurWorld = NewWorld(worldWidth, worldHeight)
		game.OurWorld.GenerateSimple()
	}

	// Check that there are sufficient spawn locations within the map.
	if noOfAIPlayers+noOfHumanPlayers > len(game.SpawnLocations) {
		return nil, fmt.Errorf("Insufficient spawn locations within map. Spawn locations found (%d)", len(game.SpawnLocations))
	}

	//fmt.Println(game.SpawnLocations)

	// Create Players
	game.Players = make([]IPlayer, noOfAIPlayers+noOfHumanPlayers)

	// Situate player bases onto the world without overlapping.
	rects, error := game.SituateHomeBases(noOfAIPlayers + noOfHumanPlayers)

	if error != nil {
		return nil, fmt.Errorf("Failed to situate home bases into world grid. Please reduce number of players and/or increase world size")
	}

	// Create Human Players
	i := 0
	for ; i < noOfHumanPlayers; i++ {
		// The world point needs to be inserted into a random location
		game.Players[i] = NewHumanPlayer(fmt.Sprintf("Human Player %d", i), rects[i].Min, playerViewWidth, playerViewHeight, game.ItemPool, game.Pathing, game.OurWorld)
		game.GenerateUnits(game.Players[i], rects[i])
	}

	// Create Machine Intelligent Players
	for j := 0; j < noOfAIPlayers; j++ {
		// The world point needs to be inserted into a random location
		game.Players[i] = NewAIPlayer(fmt.Sprintf("AI Player %d", j), rects[i].Min, playerViewWidth, playerViewHeight, game.ItemPool, game.Pathing, game.OurWorld)
		game.GenerateUnits(game.Players[i], rects[i])
		i++
	}

	// Add mechanics
	movemech := NewMovementMechanic(game.OurWorld, game.CommandChannel, game.Players, game.Pathing, &game)
	game.Mechanics = make([]IMechanic, 1)
	game.Mechanics[0] = movemech

	return &game, nil
}

// AcceptNetConnections will accept connections from UI's (humans presumably) and
// assign them a player. Once all humanplayers are accepted this method returns
// WITHOUT starting the game. We are waiting at this point ready to go.
func (game *Game) AcceptNetConnections(host string, port int) error {

	for !game.ReadyToGo() {
		// listen to incoming tcp connections
		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			return err
		}

		// Accept and if successful assign to player
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		for _, player := range game.Players {
			if player.isHuman() && !player.isWireAlive() {
				player.listen(&TCPWire{conn, json.NewDecoder(conn), json.NewEncoder(conn)})
				break
			}
		}
	}

	return nil
}

// SituateHomeBases will construct home bases in the proper
// locations on the world. That is within the world but not overlapping one another.
// It's possible for large numbers of players on a too small grid this heuristic will not converge
// and an error will be returned.
func (game *Game) SituateHomeBases(noOfPlayers int) ([]*image.Rectangle, error) {
	playerRects := make([]*image.Rectangle, noOfPlayers)

	for i := 0; i < noOfPlayers; i++ {
		// Random point within the world
		randomRect := image.Rect(game.SpawnLocations[i].Y, game.SpawnLocations[i].X, game.SpawnLocations[i].Height, game.SpawnLocations[i].Width)
		playerRects[i] = &randomRect
	}

	return playerRects, nil
}

// Start will start the game.
func (game *Game) Start() {

	// List for command on the command channel
	go game.CommandChannelHandler()

	for _, mech := range game.Mechanics {
		go mech.start()
	}

	for _, player := range game.Players {
		err := player.start()
		if err != nil {
			fmt.Printf("Player %s has failed", player.name())
			// Return this error?
		}
	}
}

// Stop will stop the game.
func (game *Game) Stop() {
	close(game.CommandChannel)

	for _, player := range game.Players {
		player.stop()
	}

	for _, mechanic := range game.Mechanics {
		mechanic.stop()
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

// FindPath finds a path from source to destination within this game's world and return it as a list of Waypoints
func (game *Game) FindPath(source *image.Point, destination *image.Point) (*list.List, error) {
	return game.Pathing.FindPath(game.ItemPool, &game.OurWorld.Grid, source, destination)
}

//FreeList will free the list return by FindPath
func (game *Game) FreeList(l *list.List) {
	game.Pathing.FreeList(game.ItemPool, l)
}

// GenerateUnits will construct the starting units per player.
func (game *Game) GenerateUnits(player IPlayer, spawnRect *image.Rectangle) {

	// Need general information about our grid and our view projection onto the grid.
	view := player.PlayerView()
	worldCenter := view.CenterOfRect(spawnRect)
	viewCenter := view.ToViewPoint(&worldCenter)

	/////////////////////////////////////////////////////////////////////////
	//           HomeStead is special. Only one in center location         //
	/////////////////////////////////////////////////////////////////////////
	homestead := HomeStead{}
	homestead.Initialize()
	homestead.generate(player)

	err := game.OurWorld.Add(&homestead, &worldCenter)
	if err != nil {
		fmt.Print(err)
	}
	homestead.CurrentLocation = &worldCenter
	player.PlayerUnits().Add(&homestead)

	/////////////////////////////////////////////////////////////////////////
	//                               All Units                             //
	/////////////////////////////////////////////////////////////////////////
	infantry := game.ItemPool.Infantry(1)
	infantry[0].generate(player)
	err = game.AddUnitCloseToPoint(player, infantry[0], &viewCenter, 10)
	if err != nil {
		fmt.Print(err)
	}

	farm := game.ItemPool.Farms(1)
	farm[0].generate(player)
	err = game.AddUnitCloseToPoint(player, farm[0], &viewCenter, 10)
	if err != nil {
		fmt.Print(err)
	}

	cavalry := game.ItemPool.Cavalry(1)
	cavalry[0].generate(player)
	err = game.AddUnitCloseToPoint(player, cavalry[0], &viewCenter, 20)
	if err != nil {
		fmt.Print(err)
	}

	woodpile := game.ItemPool.Woodpiles(1)
	woodpile[0].generate(player)
	err = game.AddUnitCloseToPoint(player, woodpile[0], &viewCenter, 30)
	if err != nil {
		fmt.Print(err)
	}

	goldmine := game.ItemPool.Goldmines(1)
	goldmine[0].generate(player)
	err = game.AddUnitCloseToPoint(player, goldmine[0], &viewCenter, 30)
	if err != nil {
		fmt.Print(err)
	}

	stonequarry := game.ItemPool.StoneQuarry(1)
	stonequarry[0].generate(player)
	err = game.AddUnitCloseToPoint(player, stonequarry[0], &viewCenter, 30)
	if err != nil {
		fmt.Print(err)
	}

}

// CommandChannelHandler will handle the command channel and dispatch
// the wire packets.
func (game *Game) CommandChannelHandler() {
	for packet := range game.CommandChannel {
		for _, player := range game.Players {
			_ = player.dispatch(packet)
		}
	}
}

// AddUnit will add a unit to this player
// without a collision within the view.
func (game *Game) AddUnit(player IPlayer, unit IUnit) {
	view := player.PlayerView()

	var locus *image.Point
	for {
		locus = view.RandomPointInView()
		if game.OurWorld.In(locus) && !game.OurWorld.Collision(locus) {
			break
		}
	}

	worldLocus := view.ToWorldPoint(locus)

	err := game.OurWorld.Add(unit, &worldLocus)
	if err != nil {
		fmt.Print(err)
	}

	unit.movement().CurrentLocation = &worldLocus

	player.PlayerUnits().Add(unit)
}

// AddUnitCloseToPoint will add unit to player no further than radius away from the central point.
// Will ensure no collisions. Central point is in VIEW coordinates.
func (game *Game) AddUnitCloseToPoint(player IPlayer, unit IUnit, central *image.Point, radius int) error {
	view := player.PlayerView()

	var locus *image.Point
	i := 0
	for {
		locus = view.RandomPointCloseToPoint(central, radius)
		if game.OurWorld.In(locus) && !game.OurWorld.Collision(locus) {
			break
		}

		// If the algorithm doesn't converge, break.
		i++
		if i > 1000 {
			return fmt.Errorf("AddUnitCloseToPoint does not converge: X(%d) Y(%d)", central.X, central.Y)
		}
	}

	worldLocus := view.ToWorldPoint(locus)

	err := game.OurWorld.Add(unit, &worldLocus)
	if err != nil {
		fmt.Print(err)
	}

	unit.movement().CurrentLocation = &worldLocus

	player.PlayerUnits().Add(unit)

	return nil
}

// LoadTMX will load the TMX (XML) file from disk (filename)
// and returns a pointer ot the tmx MAP.
// http://doc.mapeditor.org/reference/tmx-map-format/
func (game *Game) LoadTMX(filename string) (*tmx.Map, error) {

	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	m, err := tmx.Read(r)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// RenderTMX will read the TMX file (presumably already loaded) and
// render the terrain items within to our world.
func (game *Game) RenderTMX() {
	game.OurWorld = NewWorld(game.TMXMap.Height, game.TMXMap.Width)
	game.OurWorld.GenerateGrassWorld()

	// Determine the GID range of our supported terrain tiles
	for _, tileset := range game.TMXMap.Tilesets {
		switch tileset.Name {

		case "snow":
			game.SnowFirstGID = int(tileset.FirstGID)
			game.SnowLastGID = int(tileset.FirstGID) + tileset.Tilecount - 1

		case "sand":
			game.SandFirstGID = int(tileset.FirstGID)
			game.SandLastGID = int(tileset.FirstGID) + tileset.Tilecount - 1

		case "dirt":
			game.DirtFirstGID = int(tileset.FirstGID)
			game.DirtLastGID = int(tileset.FirstGID) + tileset.Tilecount - 1

		case "water":
			game.WaterFirstGID = int(tileset.FirstGID)
			game.WaterLastGID = int(tileset.FirstGID) + tileset.Tilecount - 1

		case "trees":
			game.TreesFirstGID = int(tileset.FirstGID)
			game.TreesLastGID = int(tileset.FirstGID) + tileset.Tilecount - 1

		case "grass":
			game.GrassFirstGID = int(tileset.FirstGID)
			game.GrassLastGID = int(tileset.FirstGID) + tileset.Tilecount - 1

		case "mountains":
			game.MountainsFirstGID = int(tileset.FirstGID)
			game.MountainsLastGID = int(tileset.FirstGID) + tileset.Tilecount - 1
		}

	} //for

	// Set the non-grass terrain for each layer.
	for _, layer := range game.TMXMap.Layers {
		for column := 0; column < game.TMXMap.Width; column++ {
			for row := 0; row < game.TMXMap.Height; row++ {

				switch v := int(layer.Data.DataTiles[(column + (row * game.TMXMap.Width))].GID); {

				case v >= game.MountainsFirstGID && v <= game.MountainsLastGID:
					game.OurWorld.Matrix[row][column].terrain = Mountains

				case v >= game.TreesFirstGID && v <= game.TreesLastGID:
					game.OurWorld.Matrix[row][column].terrain = Trees

				case v >= game.WaterFirstGID && v <= game.WaterLastGID:
					game.OurWorld.Matrix[row][column].terrain = Water

				case v >= game.DirtFirstGID && v <= game.DirtLastGID:
					game.OurWorld.Matrix[row][column].terrain = Dirt

				case v >= game.SandFirstGID && v <= game.SandLastGID:
					game.OurWorld.Matrix[row][column].terrain = Sand

				case v >= game.SnowFirstGID && v <= game.SnowLastGID:
					game.OurWorld.Matrix[row][column].terrain = Snow

				} //switch
			}

		}
	}

	// Search for spawn locations. These will be our suggested home base locations.
	for _, objectGroup := range game.TMXMap.ObjectGroups {
		if "spawns" == objectGroup.Name {
			// Store a *copy* of each spawn location.
			game.SpawnLocations = make([]tmx.Object, len(objectGroup.Objects))
			for i, obj := range objectGroup.Objects {

				// Spawn locations in pixels. Each acre is a 32pixel square plot.
				obj.X /= 32
				obj.Y /= 32
				obj.Width /= 32
				obj.Height /= 32

				game.SpawnLocations[i] = obj
				//fmt.Printf("%s %d %d", game.SpawnLocations[i].Name, game.SpawnLocations[i].X, game.SpawnLocations[i].Y)
			}
		}
	}

}
