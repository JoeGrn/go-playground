package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Call the function directly
	f("direct")

	// Call the function in a goroutine
	// This will run concurrently with the calling function
	go f("goroutine")

	// Starting a goroutine from an anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	// The output will be different each time you run the program
	// The goroutines are running concurrently with the main function
}
