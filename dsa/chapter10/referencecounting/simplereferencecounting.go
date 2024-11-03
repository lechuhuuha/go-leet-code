package referencecounting

import (
	"sync"
	"sync/atomic"
)

type ReferenceCounter struct {
	Num     *uint32
	pool    *sync.Pool
	Removed *uint32
	mu      sync.Mutex
	weight  int
}

func NewReferenceCounter() *ReferenceCounter {
	return &ReferenceCounter{
		Num:     new(uint32),
		pool:    &sync.Pool{},
		Removed: new(uint32),
	}
}

func (r *ReferenceCounter) Add() {
	atomic.AddUint32(r.Num, 1)

	// Automatically defer decrementing reference count
	defer r.Subtract()
}

func (r *ReferenceCounter) Subtract() {
	// Lock to ensure thread safety during modification
	r.mu.Lock()
	defer r.mu.Unlock()

	if atomic.AddUint32(r.Num, ^uint32(0)) == 0 {
		atomic.AddUint32(r.Removed, 1)
		r.cleanup()
	}
}

// cleanup handles resource cleanup when reference count is zero.
func (r *ReferenceCounter) cleanup() {
	// Resetting the pool resources or clearing objects
	r.pool.Put(nil) // Example: clear object in pool, can be customized as needed
}
