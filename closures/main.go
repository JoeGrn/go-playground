package main

// Here the variable i is made private to the int seq function and can only be updated by the exposed function
// This is a closure as the function intSeq captures the variable i

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// The variable within intSeq is maintained in memory as it is assigned to a variable within the main functions scope

func main() {
	nextInt := intSeq()
	println(nextInt())
	println(nextInt())
	println(nextInt())

	newInts := intSeq()
	println(newInts())
}
