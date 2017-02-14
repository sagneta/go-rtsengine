package rtsengine

import (
	"fmt"
	"runtime"
	"time"
)

// MovementMechanic handles the global movement of units
type MovementMechanic struct {
	// The world map that maintains the terrain and units.
	OurWorld *World

	// Command channel
	CommandChannel chan *WirePacket

	// All players in the game
	Players []IPlayer

	// Pathing systems
	Pathing *AStarPathing

	OurGame *Game

	evah bool
}

// NewMovementMechanic factory
func NewMovementMechanic(world *World, cc chan *WirePacket, players []IPlayer, pathing *AStarPathing, ourgame *Game) *MovementMechanic {
	return &MovementMechanic{world, cc, players, pathing, ourgame, true}
}

func (m *MovementMechanic) name() string {
	return "MovementMechanic"
}

func (m *MovementMechanic) stop() {
	m.evah = false
}

func (m *MovementMechanic) start() {

	for m.evah {
		for _, player := range m.Players {
			unitmap := player.PlayerUnits()
			//fmt.Printf("Go through player(%s) units ...\n", player.name())
			for _, unit := range unitmap.Map {
				movement := unit.movement()

				//fmt.Println("Found unit to move...")

				// Can/Should we move this unit?
				if movement.CanMove() {
					// As long as the MovementDestination exists and is different than the current location (obviously)
					if movement.CurrentLocation != nil && movement.MovementDestination != nil && !movement.CurrentLocation.Eq(*movement.MovementDestination) {
						//fmt.Println("Before FindPath")
						pathList, err := m.OurGame.FindPath(movement.CurrentLocation, movement.MovementDestination)
						//fmt.Println("After FindPath")
						if err != nil {
							fmt.Print(err)
							continue
						}

						// Remove the unit from the world and skip through the A* path
						m.OurWorld.RemoveAt(unit, movement.CurrentLocation)
						for e := pathList.Front(); e != nil; e = e.Next() {
							square := e.Value.(*Square)

							// Skip the current location
							if square.Locus.Eq(*movement.CurrentLocation) {
								continue
							}

							// Ok take the next point and use that and move to that point.
							// That will be our new current destination
							_ = m.OurWorld.Add(unit, &square.Locus)
							movement.CurrentLocation = &square.Locus
							break
						}

						// Free all squares to the pool.
						for e := pathList.Front(); e != nil; e = e.Next() {
							square := e.Value.(*Square)
							//square.Print()
							m.OurGame.ItemPool.Free(square)
						}

						// Update the last movement time (right now).
						unit.sendPacketToChannel(MoveUnit, m.CommandChannel)
						movement.UpdateLastMovement()
						//fmt.Print(movement.CurrentLocation)
						//fmt.Println(movement.MovementDestination)
					}
				} // move?
				runtime.Gosched()
			} // Units...
		} // Players...

		time.Sleep(time.Millisecond * 500) // half a second.
	} // eva...

}
