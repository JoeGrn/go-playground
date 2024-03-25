package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// We'll use an unsigned integer to represent our (always positive) counter.
	var ops atomic.Uint64

	// We'll use the WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		// Increment the WaitGroup counter.
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// Increment the atomic counter.
				ops.Add(1)
			}

			wg.Done()
		}()
	}
	// Wait for all goroutines to finish.
	wg.Wait()

	fmt.Println("ops:", ops.Load())
}
