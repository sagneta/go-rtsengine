package rtsengine

import (
	"log"
	"sync"
)

/*
 Maintains a pool of preallocated structures within a large array.
 The largest pool will be units. This provides two benefits.

 1) Much faster unit production
 2) Far less GC'ing that would potentially pause the game.

*/

// Pool will pool several types of structures.
type Pool struct {
	muFarms     sync.Mutex
	muArchers   sync.Mutex
	muCastles   sync.Mutex
	muCatapults sync.Mutex
	muCavalries sync.Mutex
	muFences    sync.Mutex
	muGoldmines sync.Mutex
	muInfantry  sync.Mutex
	muShips     sync.Mutex
	muQuarries  sync.Mutex
	muTowers    sync.Mutex
	muWalls     sync.Mutex
	muWoodpiles sync.Mutex

	farms     []Farm
	archers   []Archer
	castles   []Castle
	catapults []Catapult
	cavalries []Cavalry
	fences    []Fence
	goldmines []Goldmine
	infantry  []Infantry
	ships     []Ship
	quarries  []StoneQuarry
	towers    []Tower
	walls     []Wall
	woodpiles []WoodPile
}

// Generate a pool of all internal structures of maximum length
// items.
func (pool *Pool) Generate(items int) {
	pool.farms = make([]Farm, items)
	pool.archers = make([]Archer, items)
	pool.castles = make([]Castle, items)
	pool.catapults = make([]Catapult, items)
	pool.cavalries = make([]Cavalry, items)
	pool.fences = make([]Fence, items)
	pool.goldmines = make([]Goldmine, items)
	pool.infantry = make([]Infantry, items)
	pool.ships = make([]Ship, items)
	pool.quarries = make([]StoneQuarry, items)
	pool.towers = make([]Tower, items)
	pool.walls = make([]Wall, items)
	pool.woodpiles = make([]WoodPile, items)

	for i := range pool.farms {
		pool.farms[i].Deallocate()
		pool.archers[i].Deallocate()
		pool.castles[i].Deallocate()
		pool.catapults[i].Deallocate()
		pool.cavalries[i].Deallocate()
		pool.fences[i].Deallocate()
		pool.goldmines[i].Deallocate()
		pool.infantry[i].Deallocate()
		pool.ships[i].Deallocate()
		pool.quarries[i].Deallocate()
		pool.towers[i].Deallocate()
		pool.walls[i].Deallocate()
		pool.woodpiles[i].Deallocate()
	}
}

// Farms will allocate n number of farms
func (pool *Pool) Farms(n int) []*Farm {

	items := make([]*Farm, n)

	pool.muFarms.Lock()
	defer pool.muFarms.Unlock()

	j := 0
	for i := range pool.farms {
		if !pool.farms[i].IsAllocated() {
			pool.farms[i].Allocate()
			items[j] = &pool.farms[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a Farm!")
		for ; j < n; j++ {
			items[j] = &Farm{}
			items[j].Allocate()
		}
	}

	return items
}

// Free will deallocate the object and return it to the pool
func (pool *Pool) Free(object IPoolable) {
	object.Deallocate()
}
