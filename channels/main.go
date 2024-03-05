package main

import "fmt"

func main() {
	// Channels are pipes that connect concurrent goroutines
	// You can send values into channels from one goroutine and receive those values into another goroutine

	// Channels are typed by the values they convey
	messages := make(chan string)

	// Send a value into a channel using the <- syntax
	go func() { messages <- "ping" }()

	// Receive a value from a channel using the <- syntax
	msg := <-messages
	fmt.Println(msg)

	// By default, sends and receives block until both the sender and receiver are ready
	// This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization
}
