package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// If there are no messages, the default case will be executed
	// This is a non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	// If there are no messages, the default case will be executed
	// This is a non-blocking send
	// As there is no buffer or reciever
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// If there are no signals, the default case will be executed
	// This is a non-blocking receive
	// As there is no sender
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
