package rtsengine

// Poolable objects can be pooled obviously
type Poolable struct {
	allocated bool
}

// IsAllocated is true if this object was previously allocated
// and thus not available.
func (poolable *Poolable) IsAllocated() bool {
	return poolable.allocated
}

// Allocate this object from the pool
func (poolable *Poolable) Allocate() {
	poolable.allocated = true
}

// Deallocate this object from the pool
func (poolable *Poolable) Deallocate() {
	poolable.allocated = false
}
