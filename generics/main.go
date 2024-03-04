package main

import "fmt"

// MapKeys returns the keys of a map (of any type) as a slice
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// List is a generic singly linked list
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next  *element[T]
	value T
}

// We can defined methods on generic types but type parameters stay in place
// List[T] not List
func (l *List[T]) Push(v T) {
	if l.head == nil {
		l.head = &element[T]{value: v}
		l.tail = l.head
	} else {
		l.tail.next = &element[T]{value: v}
		l.tail = l.tail.next
	}
}

func (l *List[T]) GetAll() []T {
	var elems []T
	for e := l.head; e != nil; e = e.next {
		elems = append(elems, e.value)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// The type parameters are inferred when invoking generic functions
	fmt.Println("keys:", MapKeys(m))

	// But we can also specify them explicitly
	_ = MapKeys[int, string](m)

	var l = List[int]{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	fmt.Println("list:", l.GetAll())
}
