package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// create new subdirectory in the current working directory
	err := os.Mkdir("subdir", 0755)
	check(err)

	// similar to rm -rf subdir
	defer os.RemoveAll("subdir")

	// create a new empty file
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// read the contents of a directory
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")

	// iterate over the entries in the directory
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// change into a subdirectory
	err = os.Chdir("subdir/parent/child")
	check(err)

	// read the contents of the directory
	c, err = os.ReadDir(".")
	check(err)
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// change back to the original directory
	err = os.Chdir("../../..")
	check(err)

	// recursively visit a directory
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

// visit is called for each file or directory found recursively by filepath.WalkDir
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
