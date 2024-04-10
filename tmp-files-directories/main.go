package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// create a file for reading and writing, "" is default location
	f, err := os.CreateTemp("", "sample")
	check(err)

	// display name of the temp file
	// likely /tmp on unix or C:\Users\username\AppData\Local\Temp on windows
	// filename starts with given prefix
	fmt.Println("Temp file name: ", f.Name())

	// remove when done, OS will clean itself after sometime
	// but good to do this explicitly
	defer os.Remove(f.Name())

	// write some data to the file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// if you intend to write many small files, it may be more efficient to create a temporary directory
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name: ", dname)

	// clean up the directory after usage
	defer os.RemoveAll(dname)

	// create a file inside the directory
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2, 3, 4}, 0666)
	check(err)

}
