package rtsengine

import "sync"

/*
 Maintains the A* Pathing algorithm

*/

// AStarPathing will implement the A* pathing algorithm
// A simple description here: http://www.policyalmanac.org/games/aStarTutorial.htm
// Psuedocode here at the bottom of this: http://web.mit.edu/eranki/www/tutorials/search/
// https://github.com/beefsack/go-astar/blob/master/astar.go
type AStarPathing struct {
	// We need to only path-find one at a time otherwise
	// if we path-find as the world changes it will end badly.
	muPathing sync.Mutex
}
