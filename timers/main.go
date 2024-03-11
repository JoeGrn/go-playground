package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a new timer that will fire after 2 seconds
	timer1 := time.NewTimer(time.Second * 2)

	// The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// To just wait use time.Sleep
	time.Sleep(time.Second * 2)
}
