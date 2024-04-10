package main

import (
	"embed"
	"fmt"
)

// go:embed is a compiler directive
// it allows programs to include arbitary files at build time

// go:embed directive accept paths relative to the directory containing go source
// it can be used with a variable declaration to embed files

// embed a single file
// go:embed folder/single_file.txt
var fileString string

// embed content of a file into a byte
// go:embed folder/single_file.txt
var fileByte []byte

// embed multiple files with embed.FS
// creates a simple virtual file system
// go:embed folder/single_file.txt
// go:embed folder/*.hash
var folder embed.FS

func main() {
	fmt.Println(fileString)
	fmt.Println(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	fmt.Println(string(content1))
	content2, _ := folder.ReadFile("folder/file2.hash")
	fmt.Println(string(content2))
}
