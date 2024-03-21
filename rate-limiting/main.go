package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This limiter channel will receive a value every 200 milliseconds.
	limiter := time.Tick(time.Millisecond * 200)

	// by blocking on a receive from the limiter channel before serving each request.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit.
	// We can accomplish this by buffering our limiter channel.
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incoming requests.
	// The first 3 of these will benefit from the burst capability of burstyLimiter.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
