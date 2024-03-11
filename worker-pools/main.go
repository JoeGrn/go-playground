package main

import (
	"fmt"
	"time"
)

func worker(workerId int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", workerId, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", workerId, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start 3 workers
	// Initially blocekd as no jobs yet
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs and close the channel
	// As that is all the work we have
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect all the results of the work
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
