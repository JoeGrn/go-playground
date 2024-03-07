package main

import "time"

func main() {
	c1 := make(chan int, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- 1
	}()

	// Takes longer than the 1 second timeout to return
	// So the timeout case is executed
	select {
	case ch1 := <-c1:
		println("result", ch1)
	case <-time.After(1 * time.Second):
		println("timeout 1")
	}

	c2 := make(chan int, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- 2
	}()

	// Takes less than the 3 second timeout to return
	// So the result case is executed
	select {
	case ch2 := <-c2:
		println("result", ch2)
	case <-time.After(3 * time.Second):
		println("timeout 2")
	}
}
