package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// This is a simple example of encoding JSON in Go. JSON is a popular format for data exchange between systems. Go provides built-in support for encoding and decoding JSON data.
type Response1 struct {
	Page   int
	Fruits []string
}

// Fields must be exported (capitalized) to be marshalled
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// Slices and maps encode to JSON arrays and objects as you’d expect
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	// Maps will encode to JSON objects
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// The JSON package can automatically encode your custom data types.
	//It will only include exported fields in the encoded output and will by default use those names as the JSON keys
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// You can use tags on struct field declarations to customize the encoded JSON key names.
	// Check the definition of Response2 above.

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Decoding JSON data into Go values
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We need to provide a variable where the JSON package can put the decoded data.
	// This map[string]interface{} will hold a map of strings to arbitrary data types.
	var data map[string]interface{}

	// Here’s the actual decoding and a check for associated errors.
	// Unmarshall data into a pointer to a map of strings to arbitrary data types
	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}

	// Accessing nested data requires a series of conversions.
	// Here we convert the value at strs to a slice of strings.
	num := data["num"].(float64)
	fmt.Println(num)

	strs := data["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// In Go, you can also decode JSON into custom data types.
	// This has the advantages of adding additional type-safety to your programs and eliminating the need for type assertions when accessing the decoded data.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// In the examples above we always used bytes and strings as intermediates between the data and JSON representation on standard out.
	// We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
