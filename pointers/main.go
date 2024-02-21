package main

import "fmt"

// Pass by value

func zeroval(ival int) {
	ival = 0
}

// Pass by reference
// the *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address.

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// The &i syntax gives the memory address of i, i.e. a pointer to i.
	// Which is then overwriten by the zeroptr function and mutates variable i
	// Useful to prevent copying large data structures

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
