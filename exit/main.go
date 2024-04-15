package main

import (
	"fmt"
	"os"
)

func main() {
	// this will not be run
	defer fmt.Println("!")

	// immediately exit with given status
	os.Exit(3)
}
