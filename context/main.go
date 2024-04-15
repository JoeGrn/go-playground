package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// context contains information about the environment its executed in
// useful if a client disconnects or a timeout occurs
// so the server can save resource and stop serving the request

func hello(w http.ResponseWriter, r *http.Request) {
	// create a context for each request
	ctx := r.Context()
	log.Println("server: hello handler started")
	defer log.Println("server: hello handler ended")

	// wait for a few seconds before sending a response
	// to simulate a long-running process
	// select statement allows a goroutine to wait on multiple communication operations
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "Hello, World!")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	// register the handlers
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
