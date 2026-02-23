package data_structures

import (
	"errors"
	"fmt"
)

// Arena allocator uses Two Pointers to manage memory allocations
type Arena struct {
	pool  []byte // The actual memory block
	start int    // The "moving" pointer for the next allocation
	end   int    // The fixed pointer (limit)
}

// NewArena initializes the memory pool with two pointers
func NewArena(size int) *Arena {
	return &Arena{
		pool:  make([]byte, size),
		start: 0,    // Points to the beginning
		end:   size, // Points to the end (fixed safeguard)
	}
}

// Allocate moves the 'start' pointer forward
func (a *Arena) Allocate(size int) ([]byte, error) {
	// Check against the fixed 'end' pointer to prevent overflow
	if a.start+size > a.end {
		return nil, errors.New("out of memory: allocation exceeds arena bounds")
	}

	// Calculate the block to return
	allocation := a.pool[a.start : a.start+size]

	// Move the start pointer forward
	a.start += size

	return allocation, nil
}

// Deallocate moves the 'start' pointer backward
// Note: In linear allocators, you usually deallocate everything at once or in LIFO order
func (a *Arena) Deallocate(size int) {
	a.start -= size
	if a.start < 0 {
		a.start = 0
	}
}

// Reset clears the entire pool by moving 'start' back to the beginning
func (a *Arena) Reset() {
	a.start = 0
}

func main() {
	// 1. Initialize a 1KB Arena
	mem := NewArena(1024)

	// 2. Request memory (Moving 'start' forward)
	buf, _ := mem.Allocate(256)
	fmt.Printf("Allocated %d bytes. Current 'start' pointer at: %d\n", len(buf), mem.start)

	// 3. Move 'start' backward (Deallocation)
	mem.Deallocate(100)
	fmt.Printf("Deallocated 100 bytes. Current 'start' pointer at: %d\n", mem.start)
}
