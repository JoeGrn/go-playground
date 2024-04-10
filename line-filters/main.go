package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// cli tool that reads from stdin and writes to stdout capitalised
func main() {
	// Wrapping the unbuffered `os.Stdin` with a buffered scanner gives us a convenient `Scan` method
	// This advances the scanner to the next token; which is the next line in the default scanner.

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
