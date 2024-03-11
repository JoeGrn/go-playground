package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(workerId int) {
	fmt.Printf("Worker %d starting\n", workerId)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", workerId)
}

func main() {
	var wg = sync.WaitGroup{}

	// Start 5 workers
	for i := 1; i <= 5; i++ {
		// Launch go routines and increment the wait group counter
		wg.Add(1)
		// Wrap the worker in a closure
		// Which tells the wait group when the worker is done
		go func(workerId int) {
			defer wg.Done()
			worker(workerId)
		}(i)
	}

	// Block until all workers finish
	wg.Wait()

}
