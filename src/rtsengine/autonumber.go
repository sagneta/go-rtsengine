package rtsengine

import "sync"

var autonumber = 1

// AutoNumber produces and maintains an autonumber
type AutoNumber struct {
	ID           int
	muAutoNumber sync.Mutex
}

// Initialize the autonumber.
func (an *AutoNumber) Initialize() {
	an.muAutoNumber.Lock()
	defer an.muAutoNumber.Unlock()

	autonumber++
	an.ID = autonumber
}

// id returns the unique ID of this unit.
func (an *AutoNumber) id() int {
	return an.ID
}
