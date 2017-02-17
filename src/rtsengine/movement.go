package rtsengine

import (
	"container/list"
	"image"
	"time"
)

// Movement maintains state of the movement capabilities of a unit
// Stationary only units like a Home or Fence lack this structure.
type Movement struct {
	// Last actual movement of this unit
	LastMovement time.Time

	// Movement delta in milliseconds.
	// Thus if this was 500 that would be 2 movements potentially per second.
	// 1000 would be one movement per second etcetera.
	DeltaInMillis int64

	// Destination in world coordinates if a move is in progress.
	MovementDestination *image.Point

	// Current location in world coordinates
	CurrentLocation *image.Point

	// Waypath is a list of Waypoints that correspond to an entire path.
	// This allows us to calculate the paths once rather than continuously.
	Waypath *list.List
}

// CanMove returns true if this unit may move now given the current time.
// If the elapsed time Sinc Lastmovement is greater than the DeltaInMillis return true.
func (move *Movement) CanMove() bool {
	return int64(time.Since(move.LastMovement)/time.Millisecond) > move.DeltaInMillis
}

// UpdateLastMovement will update the LastMovement of this unit to the current instant in time.
func (move *Movement) UpdateLastMovement() {
	move.LastMovement = time.Now()
}
