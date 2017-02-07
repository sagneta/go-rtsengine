package rtsengine

import (
	"fmt"
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
}

func (m *MovementMechanic) name() string {
	return "MovementMechanic"
}

func (m *MovementMechanic) start() {

	for {
		for _, player := range m.Players {
			unitmap := player.PlayerUnits()
			for _, v := range unitmap.Map {
				movement := v.movement()
				if movement.CanMove() {

					if movement.CurrentLocation != nil && movement.MovementDestination != nil && !movement.CurrentLocation.Eq(*movement.MovementDestination) {
						pathList, err := m.OurGame.FindPath(movement.CurrentLocation, movement.MovementDestination)
						if err != nil {
							fmt.Print(err)
						}

						m.OurWorld.RemoveAt(v, movement.CurrentLocation)
						for e := pathList.Front(); e != nil; e = e.Next() {
							square := e.Value.(*Square)

							if square.Locus.Eq(*movement.CurrentLocation) {
								continue
							}

							m.OurWorld.Add(v, &square.Locus)
							movement.CurrentLocation = &square.Locus
						}
					}
					movement.UpdateLastMovement()
				} // move?
			} // Units...
		} // Players...

		time.Sleep(time.Millisecond * 500)
	}
}
