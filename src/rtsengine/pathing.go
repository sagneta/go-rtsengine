package rtsengine

import (
	"container/list"
	"fmt"
	"image"
)

/*
 Maintains the A* Pathing algorithm

*/

// AStarPathing will implement the A* pathing algorithm
// A simple description here: http://www.policyalmanac.org/games/aStarTutorial.htm
// Psuedocode here at the bottom of this: http://web.mit.edu/eranki/www/tutorials/search/
// https://github.com/beefsack/go-astar/blob/master/astar.go
// Smoothing to avoid diagonels: http://www.gamasutra.com/view/feature/131505/toward_more_realistic_pathfinding.php?page=1
// Basic line scan/ploting algorithm: https://en.wikipedia.org/wiki/Line_drawing_algorithm
type AStarPathing struct {
	// We need to only path-find one at a time otherwise
	// if we path-find as the world changes it will end badly.
	//muPathing sync.Mutex
}

// FindPath will find a path between source and destination Points and
// returns a list of Waypoints of the proper path.
// All coordinates in world coordinates (absolute coordinates) please.
func (path *AStarPathing) FindPath(pool *Pool, grid *Grid, source *image.Point, destination *image.Point) (*list.List, error) {

	// Check if both source and destination are not colliding
	if !grid.In(source) {
		return nil, fmt.Errorf("Source not in grid! (%d,%d)", source.X, source.Y)
	}

	/*
		if grid.Collision(source) {
			return nil, fmt.Errorf("Source collision! (%d,%d)", source.X, source.Y)
		}
	*/

	if !grid.In(destination) || grid.Collision(destination) {
		return nil, fmt.Errorf("Destination not in grid or collision! (%d,%d)", destination.X, destination.Y)
	}

	closedList := list.New()
	openList := list.New()

	// Starting waypoint. 0 out the cost.
	q := pool.Waypoints(1)[0]
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
		//find the waypoint with the least f on the open list, call it "q"
		//remove q from the open list
		q = path.leastF(openList)

		// generate q's 8 successors and set their parents to q
		successors := path.constructSuccessor(pool, q)
		for i, successor := range successors {

			// ensure it is in the grid and there isn't a collision
			if !grid.In(&successor.Locus) || grid.Collision(&successor.Locus) {
				pool.Free(successor)
				continue
			}

			//if successor is the goal, stop the search
			if destination.Eq(successor.Locus) {
				closedList.PushBack(q)
				closedList.PushBack(successor)
				path.FreeList(pool, openList)
				path.freeArray(pool, i+1, successors)

				return path.smoothPath(grid, pool, closedList), nil
			}

			//successor.g = q.g + distance between successor and q
			D := grid.DistanceDiagonelShortcut(&q.Locus, &successor.Locus)
			successor.G = q.G + D

			// successor.h = distance from goal to successor
			successor.H = grid.DistanceDiagonelShortcut(&successor.Locus, destination)

			// successor.f = successor.g + successor.h
			successor.F = successor.G + successor.H

			//  if a waypoint with the same position as successor is in the OPEN list
			//  exists and has a lower f than successor, skip this successor
			if path.skipSuccessor(successor, openList) {
				pool.Free(successor)
				continue
			}

			// if a waypoint with the same position as successor is in the CLOSED list
			// exists has a lower f than successor, skip this successor
			if path.skipSuccessor(successor, closedList) {
				pool.Free(successor)
				continue
			}

			// otherwise, add the waypoint to the open list
			openList.PushBack(successor)

		} // for successors

		// push q on the closed list
		closedList.PushBack(q)

	} // openList non empty

	// Free all the remaining successors in the open list.
	path.FreeList(pool, openList)

	return path.smoothPath(grid, pool, closedList), nil
}

// freeArray will free all waypoints in array from i .. len(waypoints)-1
func (path *AStarPathing) freeArray(pool *Pool, i int, waypoints []*Waypoint) {
	if i >= len(waypoints) {
		return
	}

	for ; i < len(waypoints); i++ {
		pool.Free(waypoints[i])
	}
}

// FreeList will free every Waypoint in the list l
func (path *AStarPathing) FreeList(pool *Pool, l *list.List) {
	// Free all the remaining successors in the open list.
	for e := l.Front(); e != nil; e = e.Next() {
		pool.Free(e.Value.(*Waypoint))
	}
}

// smoothPath will smooth the path making it more direct using a variation walkable algorithm below:
// http://www.gamasutra.com/view/feature/131505/toward_more_realistic_pathfinding.php?page=2
func (path *AStarPathing) smoothPath(grid *Grid, pool *Pool, l *list.List) *list.List {
	result := list.New()
	l = path.optimizePath(pool, l)

	checkPoint := l.Front().Value.(*Waypoint)
	for e := l.Front().Next(); e != nil; e = e.Next() {
		currentPoint := e.Value.(*Waypoint)
		_, ok := path.walkable(grid, checkPoint, currentPoint)
		if ok {
			continue
		}

		result.PushBack(checkPoint)
		checkPoint = currentPoint
	}
	result.PushBack(checkPoint) // Don't forget the last point

	// Now we have a sparse list of waypoints with gaps between them.
	// We fill in the gaps with straight lines.
	smoothPath := list.New()

	from := result.Front().Value.(*Waypoint)
	s1 := pool.Waypoints(1)[0]
	s1.Locus = from.Locus
	smoothPath.PushBack(s1)
	for e := result.Front().Next(); e != nil; e = e.Next() {
		to := e.Value.(*Waypoint)
		points := grid.DirectLineBresenham(&from.Locus, &to.Locus)

		points = points[1:] // Cull the source

		waypoints := pool.Waypoints(len(points))
		for i, point := range points {
			waypoints[i].Locus = point
			smoothPath.PushBack(waypoints[i])
		}
	}

	// Free them all. You are done with them.
	path.FreeList(pool, result)

	return smoothPath
}

// walkable returns all the points between from and to if it is walkable (no collisions) with a true boolean.
// If no such direct walkable path exists then nil, false is returned.
func (path *AStarPathing) walkable(grid *Grid, from *Waypoint, to *Waypoint) ([]image.Point, bool) {
	points := grid.DirectLineBresenham(&from.Locus, &to.Locus)

	for _, point := range points {
		if grid.Collision(&point) {
			return nil, false
		}
	}

	return points, true
}

// optimizePath will optimize the path list passed as a parameter. Any culled
// waypoints are freed from the pool.
//
// A path list will contain duplicates at each _position_. Thus you want to
// iterate over the list and remove duplicates at each _position_ leaving the
// waypoint with the least F in the path list.
// For F ties only one is chosen.
func (path *AStarPathing) optimizePath(pool *Pool, l *list.List) *list.List {
	var m map[int]*Waypoint

	m = make(map[int]*Waypoint)
	for e := l.Front(); e != nil; e = e.Next() {
		waypoint := e.Value.(*Waypoint)

		p, ok := m[waypoint.Position]

		if !ok {
			m[waypoint.Position] = waypoint
		} else {
			if p.F <= waypoint.F {
				pool.Free(waypoint)
			} else {
				m[waypoint.Position] = waypoint
				pool.Free(p)
			}
		}

	}
	result := list.New()

	length := len(m)
	for i := 0; i < length; i++ {
		result.PushBack(m[i])
	}

	return result
}

//

// skipSuccessor will scan list l and if the list l contains an element with a smaller
// F than the successor at the same position, returns TRUE.
func (path *AStarPathing) skipSuccessor(successor *Waypoint, l *list.List) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		waypoint := e.Value.(*Waypoint)

		if waypoint.Position == successor.Position && waypoint.F <= successor.F {
			return true
		}
	}

	return false
}

// constructSuccessor will construct 8 successors with parent q from the Pool pool.
func (path *AStarPathing) constructSuccessor(pool *Pool, q *Waypoint) []*Waypoint {
	// The successors are the adjoining waypoints with the source S
	// in the middle. See below. It moves clockwise. We index
	// from zero so index 0 is 1 below.
	// 1  2  3
	// 8  S  4
	// 7  6  5
	successors := pool.Waypoints(8)

	// 1
	successors[0].Locus.X = q.Locus.X - 1
	successors[0].Locus.Y = q.Locus.Y - 1
	successors[0].Parent = q
	successors[0].Position = q.Position + 1
	successors[0].F = 0
	successors[0].G = 0
	successors[0].H = 0

	// 2
	successors[1].Locus.X = q.Locus.X - 1
	successors[1].Parent = q
	successors[1].Position = q.Position + 1
	successors[1].F = 0
	successors[1].G = 0
	successors[1].H = 0

	// 3
	successors[2].Locus.X = q.Locus.X - 1
	successors[2].Locus.Y = q.Locus.Y + 1
	successors[2].Parent = q
	successors[2].Position = q.Position + 1
	successors[2].F = 0
	successors[2].G = 0
	successors[2].H = 0

	// 4
	successors[3].Locus.Y = q.Locus.Y + 1
	successors[3].Parent = q
	successors[3].Position = q.Position + 1
	successors[3].F = 0
	successors[3].G = 0
	successors[3].H = 0

	// 5
	successors[4].Locus.X = q.Locus.X + 1
	successors[4].Locus.Y = q.Locus.Y + 1
	successors[4].Parent = q
	successors[4].Position = q.Position + 1
	successors[4].F = 0
	successors[4].G = 0
	successors[4].H = 0

	// 6
	successors[5].Locus.X = q.Locus.X + 1
	successors[5].Parent = q
	successors[5].Position = q.Position + 1
	successors[5].F = 0
	successors[5].G = 0
	successors[5].H = 0

	// 7
	successors[6].Locus.X = q.Locus.X + 1
	successors[6].Locus.Y = q.Locus.Y - 1
	successors[6].Parent = q
	successors[6].Position = q.Position + 1
	successors[6].F = 0
	successors[6].G = 0
	successors[6].H = 0

	// 8
	successors[7].Locus.Y = q.Locus.Y - 1
	successors[7].Parent = q
	successors[7].Position = q.Position + 1
	successors[7].F = 0
	successors[7].G = 0
	successors[7].H = 0

	return successors
}

// leastF returns the Waypoint with the least F within list l
// AND remove that Waypoint from list l.
// Returns nil if no item exists.
func (path *AStarPathing) leastF(l *list.List) *Waypoint {

	var leastWaypoint *Waypoint
	var leastWaypointE *list.Element
	for e := l.Front(); e != nil; e = e.Next() {
		waypoint := e.Value.(*Waypoint)
		if leastWaypoint == nil || waypoint.F < leastWaypoint.F {
			leastWaypoint = waypoint
			leastWaypointE = e
		}
	}

	l.Remove(leastWaypointE)
	return leastWaypoint
}
