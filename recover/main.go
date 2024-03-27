package main

import "fmt"

func mayPanic() {
	panic("an error occurred")
}

// Recover is a built-in function that regains control of a panicking goroutine.
// An example use of recover is to shut down a failing goroutine inside a server without crashing the whole server.

func main() {
	defer func() {
		// Recover is only useful inside deferred functions.
		// The return value is the error value that was passed to the call to panic.
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("Exiting normally")
}
