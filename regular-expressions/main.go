package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// match a string
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	// compile an optimised regexp struct
	r, _ := regexp.Compile("a([a-z]+)e")

	// match a string
	fmt.Println(r.MatchString("apple"))

	// find the first match
	fmt.Println(r.FindString("apple tree"))

	// find the first match and return the start and end indexes
	fmt.Println("index: ", r.FindStringIndex("apple tree"))

	// find whole patten and sub pattern matches
	fmt.Println("index: ", r.FindStringSubmatch("apple tree"))

	// find indexes and sub pattern indexes
	fmt.Println("index: ", r.FindStringSubmatchIndex("apple tree"))

	// find all matches
	fmt.Println(r.FindAllString("apple tree apple tree", -1))

	// find all index matches not just the first
	fmt.Println("all: ", r.FindAllStringSubmatch("apple tree apple tree", -1))

	// provide a limit to the number of matches
	fmt.Println(r.FindAllString("apple tree apple tree", 2))

	// can also provide byte slices
	fmt.Println(r.Match([]byte("apple")))

	// replace all matches
	fmt.Println(r.ReplaceAllString("apple tree apple tree", "<fruit>"))

	// replace all matches with a function
	in := []byte("a apple")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
