package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// it is idiomatic to encapsulate struct creation in a function
	fmt.Println(newPerson("Jon"))

	// access struct fields with dot notation
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// pointers to structs
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// structs for single values can be anonymous
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
