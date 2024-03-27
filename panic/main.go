package main

import "os"

func main() {
	panic("an error occurred")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
