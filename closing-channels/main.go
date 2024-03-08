package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// This is a worker go routine
	// It will receive jobs from the jobs channel
	// The second return value will be false if jobs has been closed and all values in the channel have already been received
	// We'll use this to notify on done when we've worked all our jobs
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// This sends 3 jobs to the worker over the jobs channel, then closes it
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// Await the worker using the synchronization approach
	<-done

	// Reading from a closed channel will return the zero value immediately
	// Test that all jobs have been received

	_, ok := <-jobs
	fmt.Println("more jobs:", ok)
}
