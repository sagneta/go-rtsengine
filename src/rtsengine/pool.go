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

 Should a request for an unit fail to find a free structure
 within the pool then one will be dyanamically created.
 Upon invocation of Free that structure will be GC'ed like
 any other lacking a reference.
*/

// Pool will pool several types of structures.
type Pool struct {
	muFarms     sync.Mutex
	muArchers   sync.Mutex
	muCastles   sync.Mutex
	muCatapults sync.Mutex
	muCavalry   sync.Mutex
	muFences    sync.Mutex
	muGoldmines sync.Mutex
	muInfantry  sync.Mutex
	muShips     sync.Mutex
	muQuarries  sync.Mutex
	muTowers    sync.Mutex
	muWalls     sync.Mutex
	muWoodpiles sync.Mutex
	muPeasants  sync.Mutex

	farms     []Farm
	archers   []Archer
	castles   []Castle
	catapults []Catapult
	cavalry   []Cavalry
	fences    []Fence
	goldmines []Goldmine
	infantry  []Infantry
	ships     []Ship
	quarries  []StoneQuarry
	towers    []Tower
	walls     []Wall
	woodpiles []WoodPile
	peasants  []Peasant
}

// Generate a pool of all internal structures of maximum length
// items.
func (pool *Pool) Generate(items int) {
	pool.farms = make([]Farm, items)
	pool.archers = make([]Archer, items)
	pool.castles = make([]Castle, items)
	pool.catapults = make([]Catapult, items)
	pool.cavalry = make([]Cavalry, items)
	pool.fences = make([]Fence, items)
	pool.goldmines = make([]Goldmine, items)
	pool.infantry = make([]Infantry, items)
	pool.ships = make([]Ship, items)
	pool.quarries = make([]StoneQuarry, items)
	pool.towers = make([]Tower, items)
	pool.walls = make([]Wall, items)
	pool.woodpiles = make([]WoodPile, items)
	pool.peasants = make([]Peasant, items)

	for i := range pool.farms {
		pool.farms[i].Deallocate()
		pool.archers[i].Deallocate()
		pool.castles[i].Deallocate()
		pool.catapults[i].Deallocate()
		pool.cavalry[i].Deallocate()
		pool.fences[i].Deallocate()
		pool.goldmines[i].Deallocate()
		pool.infantry[i].Deallocate()
		pool.ships[i].Deallocate()
		pool.quarries[i].Deallocate()
		pool.towers[i].Deallocate()
		pool.walls[i].Deallocate()
		pool.woodpiles[i].Deallocate()
		pool.peasants[i].Deallocate()
	}
}

// Farms allocated n at a time.
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

// Cavalry allocated n at a time.
func (pool *Pool) Cavalry(n int) []*Cavalry {

	items := make([]*Cavalry, n)

	pool.muCavalry.Lock()
	defer pool.muCavalry.Unlock()

	j := 0
	for i := range pool.cavalry {
		if !pool.cavalry[i].IsAllocated() {
			pool.cavalry[i].Allocate()
			items[j] = &pool.cavalry[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a Cavalry!")
		for ; j < n; j++ {
			items[j] = &Cavalry{}
			items[j].Allocate()
		}
	}

	return items
}

// StoneQuarry allocated n at a time.
func (pool *Pool) StoneQuarry(n int) []*StoneQuarry {

	items := make([]*StoneQuarry, n)

	pool.muQuarries.Lock()
	defer pool.muQuarries.Unlock()

	j := 0
	for i := range pool.quarries {
		if !pool.quarries[i].IsAllocated() {
			pool.quarries[i].Allocate()
			items[j] = &pool.quarries[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a StoneQuarry!")
		for ; j < n; j++ {
			items[j] = &StoneQuarry{}
			items[j].Allocate()
		}
	}

	return items
}

// Infantry allocated n at a time.
func (pool *Pool) Infantry(n int) []*Infantry {

	items := make([]*Infantry, n)

	pool.muInfantry.Lock()
	defer pool.muInfantry.Unlock()

	j := 0
	for i := range pool.infantry {
		if !pool.infantry[i].IsAllocated() {
			pool.infantry[i].Allocate()
			items[j] = &pool.infantry[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a Infantry!")
		for ; j < n; j++ {
			items[j] = &Infantry{}
			items[j].Allocate()
		}
	}

	return items
}

// Archers allocated n at a time.
func (pool *Pool) Archers(n int) []*Archer {

	items := make([]*Archer, n)

	pool.muArchers.Lock()
	defer pool.muArchers.Unlock()

	j := 0
	for i := range pool.archers {
		if !pool.archers[i].IsAllocated() {
			pool.archers[i].Allocate()
			items[j] = &pool.archers[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a Archer!")
		for ; j < n; j++ {
			items[j] = &Archer{}
			items[j].Allocate()
		}
	}

	return items
}

// Castles allocated n at a time.
func (pool *Pool) Castles(n int) []*Castle {

	items := make([]*Castle, n)

	pool.muCastles.Lock()
	defer pool.muCastles.Unlock()

	j := 0
	for i := range pool.castles {
		if !pool.castles[i].IsAllocated() {
			pool.castles[i].Allocate()
			items[j] = &pool.castles[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a Castle!")
		for ; j < n; j++ {
			items[j] = &Castle{}
			items[j].Allocate()
		}
	}

	return items
}

// Catapults allocated n at a time.
func (pool *Pool) Catapults(n int) []*Catapult {

	items := make([]*Catapult, n)

	pool.muCatapults.Lock()
	defer pool.muCatapults.Unlock()

	j := 0
	for i := range pool.catapults {
		if !pool.catapults[i].IsAllocated() {
			pool.catapults[i].Allocate()
			items[j] = &pool.catapults[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a Catapult!")
		for ; j < n; j++ {
			items[j] = &Catapult{}
			items[j].Allocate()
		}
	}

	return items
}

// Fences allocated n at a time.
func (pool *Pool) Fences(n int) []*Fence {

	items := make([]*Fence, n)

	pool.muFences.Lock()
	defer pool.muFences.Unlock()

	j := 0
	for i := range pool.fences {
		if !pool.fences[i].IsAllocated() {
			pool.fences[i].Allocate()
			items[j] = &pool.fences[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a Fence!")
		for ; j < n; j++ {
			items[j] = &Fence{}
			items[j].Allocate()
		}
	}

	return items
}

// Goldmines allocated n at a time.
func (pool *Pool) Goldmines(n int) []*Goldmine {

	items := make([]*Goldmine, n)

	pool.muGoldmines.Lock()
	defer pool.muGoldmines.Unlock()

	j := 0
	for i := range pool.goldmines {
		if !pool.goldmines[i].IsAllocated() {
			pool.goldmines[i].Allocate()
			items[j] = &pool.goldmines[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a Goldmine!")
		for ; j < n; j++ {
			items[j] = &Goldmine{}
			items[j].Allocate()
		}
	}

	return items
}

// Ships allocated n at a time.
func (pool *Pool) Ships(n int) []*Ship {

	items := make([]*Ship, n)

	pool.muShips.Lock()
	defer pool.muShips.Unlock()

	j := 0
	for i := range pool.ships {
		if !pool.ships[i].IsAllocated() {
			pool.ships[i].Allocate()
			items[j] = &pool.ships[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a Ship!")
		for ; j < n; j++ {
			items[j] = &Ship{}
			items[j].Allocate()
		}
	}

	return items
}

// Towers allocated n at a time.
func (pool *Pool) Towers(n int) []*Tower {

	items := make([]*Tower, n)

	pool.muTowers.Lock()
	defer pool.muTowers.Unlock()

	j := 0
	for i := range pool.towers {
		if !pool.towers[i].IsAllocated() {
			pool.towers[i].Allocate()
			items[j] = &pool.towers[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a Tower!")
		for ; j < n; j++ {
			items[j] = &Tower{}
			items[j].Allocate()
		}
	}

	return items
}

// Walls allocated n at a time.
func (pool *Pool) Walls(n int) []*Wall {

	items := make([]*Wall, n)

	pool.muWalls.Lock()
	defer pool.muWalls.Unlock()

	j := 0
	for i := range pool.walls {
		if !pool.walls[i].IsAllocated() {
			pool.walls[i].Allocate()
			items[j] = &pool.walls[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a Wall!")
		for ; j < n; j++ {
			items[j] = &Wall{}
			items[j].Allocate()
		}
	}

	return items
}

// Woodpiles allocated n at a time.
func (pool *Pool) Woodpiles(n int) []*WoodPile {

	items := make([]*WoodPile, n)

	pool.muWoodpiles.Lock()
	defer pool.muWoodpiles.Unlock()

	j := 0
	for i := range pool.woodpiles {
		if !pool.woodpiles[i].IsAllocated() {
			pool.woodpiles[i].Allocate()
			items[j] = &pool.woodpiles[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a WoodPile!")
		for ; j < n; j++ {
			items[j] = &WoodPile{}
			items[j].Allocate()
		}
	}

	return items
}

// Peasants allocated n at a time.
func (pool *Pool) Peasants(n int) []*Peasant {

	items := make([]*Peasant, n)

	pool.muPeasants.Lock()
	defer pool.muPeasants.Unlock()

	j := 0
	for i := range pool.peasants {
		if !pool.peasants[i].IsAllocated() {
			pool.peasants[i].Allocate()
			items[j] = &pool.peasants[i]
			j++
		}
		if j >= n {
			break
		}
	}

	if j < n {
		log.Print("We failed to allocate a Peasant!")
		for ; j < n; j++ {
			items[j] = &Peasant{}
			items[j].Allocate()
		}
	}

	return items
}

// Free will deallocate the object and return it to the pool
func (pool *Pool) Free(objects ...IPoolable) {
	for _, object := range objects {
		object.Deallocate()
	}
}
