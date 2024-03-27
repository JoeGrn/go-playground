package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// Here's an example that sorts a slice of strings by their length.
	// Returns -1 if a < b, 0 if a == b, and 1 if a > b.
	lenCompare := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Sort Func uses the given comparison function to sort the slice.
	slices.SortFunc(fruits, lenCompare)
	fmt.Println(fruits)

	// Can sort values which are not buult in types

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Bob", age: 31},
		Person{name: "John", age: 42},
		Person{name: "Michael", age: 17},
	}

	// Sort by age.
	// SortFunc take an anonymous function

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)
}
