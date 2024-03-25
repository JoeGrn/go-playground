package main

import (
	"fmt"
	"sync"
)

type container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *container) increment(key string) {
	// Lock so only one goroutine at a time can access the map c.counters
	c.mu.Lock()
	// Unlock when the function is done using a defer statement
	defer c.mu.Unlock()

	c.counters[key]++
}

func main() {
	c := container{
		counters: make(map[string]int),
	}
	var wg sync.WaitGroup

	doIncrement := func(key string, n int) {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			c.increment(key)
		}
	}
	wg.Add(3)
	go doIncrement("a", 1000)
	go doIncrement("b", 1000)
	go doIncrement("b", 1000)

	wg.Wait()
	fmt.Println(c.counters)
}
