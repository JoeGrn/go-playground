package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Issue an HTTP GET request to a server.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("Response status:", resp.Status)

	// Print the response body.
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
