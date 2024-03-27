package main

import (
	"fmt"
	"slices"
)

// Sorting functions are generic, and work for any ordered built-in type. For a list of ordered types, see cmp.Ordered.

func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println(strs)

	ints := []int{3, 1, 2}
	slices.Sort(ints)
	fmt.Println(ints)

	s := slices.IsSorted(ints)
	fmt.Println(s)
}
