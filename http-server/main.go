package main

import (
	"fmt"
	"net/http"
)

// a handler for the /hello route
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// a handler for the /headers route
// this handler prints all the headers in the request
func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// register the handlers
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// start the server
	// nil means to use the default server mux
	http.ListenAndServe(":8090", nil)
}
