package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	fmt.Println(now.UnixMilli())

	// convert to time
	fmt.Println((time.Unix(now.Unix(), 0)))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
