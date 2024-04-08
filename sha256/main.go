package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// The string to hash.
	s := "sha256"

	// New returns a new hash.Hash computing the SHA256 checksum.
	h := sha256.New()

	// Write (via the embedded io.Writer interface) adds more data to the running hash.
	h.Write([]byte(s))

	// Sum appends the current hash to b and returns the resulting slice.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
