package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	// Prints an instance of the point struct
	fmt.Printf("%v\n", p)

	// Prints the struct's field names
	fmt.Printf("%+v\n", p)

	// Prints the struct's field names and types
	fmt.Printf("%#v\n", p)

	// Prints the struct's types
	fmt.Printf("%T\n", p)

	// Prints a boolean
	fmt.Printf("%t\n", true)

	// Prints a number in base 10
	fmt.Printf("%d\n", 123)

	// Prints a number in binary
	fmt.Printf("%b\n", 14)

	// Prints a character that corresponds to the given integer
	fmt.Printf("%c\n", 33)

	// Prints a number in hex
	fmt.Printf("%x\n", 456)

	// Prints a number in float
	fmt.Printf("%f\n", 78.9)

	// Prints a number in scientific notation
	fmt.Printf("%e\n", 123400000.0)

	// Prints a string
	fmt.Printf("%s\n", "\"string\"")

	// Prints a string in double quotes
	fmt.Printf("%q\n", "\"string\"")

	// Prints a string in base 16
	fmt.Printf("%x\n", "hex this")

	// Prints a pointer
	fmt.Printf("%p\n", &p)

	// Prints a number in width 6, right-justified
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// Prints a number in width 6, left-justified
	fmt.Printf("|%-6d|%-6d|\n", 12, 345)

	// Prints a number in width 6, with 2 decimal places
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// Prints a string in width 6, right-justified
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// Prints a string in width 6, left-justified
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// Prints a string in width 6, with 2 decimal places
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Prints to the io.Writer
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
