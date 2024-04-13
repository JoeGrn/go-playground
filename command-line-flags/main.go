package main

// $ go build command-line-flags.go
// $ ./command-line-flags -word=opt -numb=7 -fork -svar=flag

import (
	"flag"
	"fmt"
)

// go provides a flag package supporting basic command-line flag parsing
// if you omit flags they automatically get their default values
// the flag package requires all flags to appear before positional arguments
// -h and --help are automatically provided by the flag package
// if you provide an undefined flag, the program will show an error message and the help text

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// once all flags are declared, call flag.Parse() to execute the command-line parsing
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
