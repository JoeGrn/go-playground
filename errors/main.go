package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if (arg) == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// Custom error type
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// Return a pointer to the custom error type
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// If you want to use the data in a custom error, you'll need to get the error as an instance of the custom error type via type assertion
	// ok is whether the type assetion was successful
	_, e := f2(42)
	if argError, ok := e.(*argError); ok {
		fmt.Println(argError.arg)
		fmt.Println(argError.prob)
	}
}
