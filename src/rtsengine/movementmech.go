package rtsengine

import (
	"fmt"
	"image"
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
			for _, unit := range unitmap.Map {
				movement := unit.movement()

				// Can/Should we move this unit?
				if movement.CanMove() {
					// As long as the MovementDestination exists and is different than the current location (obviously)
					if movement.CurrentLocation != nil && movement.MovementDestination != nil && !movement.CurrentLocation.Eq(*movement.MovementDestination) {
						//fmt.Printf("Move Unit(%d) in movementmech to %s. \n", unit.id(), movement.MovementDestination)
						pathList, err := m.OurGame.FindPath(movement.CurrentLocation, movement.MovementDestination)
						if err != nil {
							movement.MovementDestination = nil // stop further movement
							fmt.Print(err)
							continue
						}

						// Remove the unit from the world and skip through the A* path
						m.OurWorld.RemoveAt(unit, movement.CurrentLocation)
						for e := pathList.Front(); e != nil; e = e.Next() {
							waypoint := e.Value.(*Waypoint)

							// Skip the current location
							if waypoint.Locus.Eq(*movement.CurrentLocation) {
								continue
							}

							// Ok take the next point and use that and move to that point.
							// That will be our new current destination.
							// Ensure to use a copy of the waypoint as the waypoint is pooled.
							newLocation := image.Pt(waypoint.Locus.X, waypoint.Locus.Y)
							_ = m.OurWorld.Add(unit, &newLocation)
							movement.CurrentLocation = &newLocation

							// Remove the head of the pathList as that will be the current location.
							// Insert the calculated pathList as the current waypath
							break
						}

						// Update the last movement time (right now).
						unit.sendPacketToChannel(MoveUnit, m.CommandChannel)
						movement.UpdateLastMovement()

						// Free all waypoints to the pool.
						m.OurGame.FreeList(pathList)
					}
				} // move?
				runtime.Gosched()
			} // Units...
		} // Players...

		time.Sleep(time.Millisecond * 500) // half a second.
	} // eva...

}
