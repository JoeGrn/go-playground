package main

// $ go build command-line-arguments.go
// $ ./command-line-arguments a b c d

import (
	"fmt"
	"os"
)

func main() {
	// raw cli args
	argsWithProg := os.Args

	// the first value in this slice is the path to the program
	// os.Args[1:] holds the arguments to the program
	argsWithoutProg := os.Args[1:]

	// get the value of a specific argument
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
