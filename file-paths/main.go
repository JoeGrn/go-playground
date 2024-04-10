package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// contruct a path
	p := filepath.Join("dir1", "dir2", "file")
	fmt.Println("p:", p)

	// use join instead of string concatenation
	// join also normalizes paths by removing superfluous separators and directory changes
	fmt.Println(filepath.Join("dir1//", "file"))
	fmt.Println(filepath.Join("dir1/../dir1", "file"))

	// Dir and Base can be used to split a path to the directory and file
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// Check if a path is absolute
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	// Get the extension of a file
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// To find the file name with the extension removed, use strings.TrimSuffix
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel finds a relative path between a base and target
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
