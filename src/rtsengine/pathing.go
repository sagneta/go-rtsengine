package rtsengine

import (
	"container/list"
	"fmt"
	"image"
	"sync"
)

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

// FindPath will find a path between source and destination Points and
// returns a list of Squares of the proper path.
// All coordinates in world coordinates (absolute coordinates) please.
func (path *AStarPathing) FindPath(pool *Pool, grid *Grid, source *image.Point, destination *image.Point) (*list.List, error) {

	// Check if both source and destination are not colliding
	if !grid.In(source) || grid.Collision(source) {
		return nil, fmt.Errorf("Source not in grid or collision! (%d,%d)", source.X, source.Y)
	}

	if !grid.In(destination) || grid.Collision(destination) {
		return nil, fmt.Errorf("Destination not in grid or collision! (%d,%d)", destination.X, destination.Y)
	}

	closedList := list.New()
	openList := list.New()

	// Starting square. 0 out the cost.
	q := pool.Squares(1)[0]
	q.F = 0
	q.G = 0
	q.H = 0
	q.Locus.X = source.X
	q.Locus.Y = source.Y
	q.Position = 0

	// Push onto the openlist to prime the pathing engine
	openList.PushFront(q)

	// While the open list is not empty
	for openList.Len() > 0 {
		//find the node with the least f on the open list, call it "q"
		//remove q from the open list
		q = path.leastF(openList)

		// generate q's 8 successors and set their parents to q
		successors := path.constructSuccessor(pool, q)
		for _, successor := range successors {

			// ensure it is in the grid and there isn't a collision
			if !grid.In(&successor.Locus) || grid.Collision(&successor.Locus) {
				pool.Free(successor)
				continue
			}

			//if successor is the goal, stop the search
			if destination.Eq(successor.Locus) {
				closedList.PushBack(q)
				closedList.PushBack(successor)
				return closedList, nil
			}

			//successor.g = q.g + distance between successor and q
			D := grid.Distance(&q.Locus, &successor.Locus)
			successor.G = q.G + D

			// successor.h = distance from goal to successor
			successor.H = grid.Distance(&successor.Locus, destination)

			// successor.f = successor.g + successor.h
			successor.F = successor.G + successor.H

			//  if a node with the same position as successor is in the OPEN list
			//  which has a lower f than successor, skip this successor
			if path.skipSuccessor(successor, openList) {
				pool.Free(successor)
				continue
			}

			// if a node with the same position as successor is in the CLOSED list \
			// which has a lower f than successor, skip this successor
			if path.skipSuccessor(successor, closedList) {
				pool.Free(successor)
				continue
			}

			// otherwise, add the node to the open list
			openList.PushBack(successor)

		} // for successors

		// push q on the closed list
		closedList.PushBack(q)

	} // openList non empty

	// Free all the remaining successors in the open list.
	for e := openList.Front(); e != nil; e = e.Next() {
		square := e.Value.(*Square)
		pool.Free(square)
	}

	return closedList, nil
}

//func (path *AStarPathing
//var m map[string]int

// skipSuccessor will scan list l and if the list l contains an element with a smaller
// F than the successor at the same position, returns TRUE.
func (path *AStarPathing) skipSuccessor(successor *Square, l *list.List) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		square := e.Value.(*Square)

		if square.Position == successor.Position && square.F <= successor.F {
			return true
		}
	}

	return false
}

// constructSuccessor will construct 8 successors with parent q from the Pool pool.
func (path *AStarPathing) constructSuccessor(pool *Pool, q *Square) []*Square {
	// The successors are the adjoining squares with the source S
	// in the middle. See below. It moves clockwise. We index
	// from zero so index 0 is 1 below.
	// 1  2  3
	// 8  S  4
	// 7  6  5
	successors := pool.Squares(8)

	// 1
	successors[0].Locus.X = q.Locus.X - 1
	successors[0].Locus.Y = q.Locus.Y - 1
	successors[0].Parent = q
	successors[0].Position = q.Position + 1

	// 2
	successors[1].Locus.X = q.Locus.X - 1
	successors[1].Parent = q
	successors[1].Position = q.Position + 1

	// 3
	successors[2].Locus.X = q.Locus.X - 1
	successors[2].Locus.Y = q.Locus.Y + 1
	successors[2].Parent = q
	successors[2].Position = q.Position + 1

	// 4
	successors[3].Locus.Y = q.Locus.Y + 1
	successors[3].Parent = q
	successors[3].Position = q.Position + 1

	// 5
	successors[4].Locus.X = q.Locus.X + 1
	successors[4].Locus.Y = q.Locus.Y + 1
	successors[4].Parent = q
	successors[4].Position = q.Position + 1

	// 6
	successors[5].Locus.X = q.Locus.X + 1
	successors[5].Parent = q
	successors[5].Position = q.Position + 1

	// 7
	successors[6].Locus.X = q.Locus.X + 1
	successors[6].Locus.Y = q.Locus.Y - 1
	successors[6].Parent = q
	successors[6].Position = q.Position + 1

	// 8
	successors[7].Locus.Y = q.Locus.Y - 1
	successors[7].Parent = q
	successors[7].Position = q.Position + 1

	return successors
}

// leastF returns the Square with the least F within list l
// AND remove that Square from list l.
// Returns nil if no item exists.
func (path *AStarPathing) leastF(l *list.List) *Square {

	var leastSquare *Square
	var leastSquareE *list.Element
	for e := l.Front(); e != nil; e = e.Next() {
		square := e.Value.(*Square)
		if leastSquare == nil || square.F < leastSquare.F {
			leastSquare = square
			leastSquareE = e
		}
	}

	l.Remove(leastSquareE)
	return leastSquare
}
