package rtsengine

// IPoolable objects can be pooled obviously
type IPoolable interface {
	IsAllocated() bool
	Allocate()
	Deallocate()
}
