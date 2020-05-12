/*
Easily write a basic HTTP server
*/ 

package main

import (
	"fmt"
	"net/http"
)

// HANDLERS
// a handler is an object implementing the http.Handler interface
func hello(w http.ResponseWriter, req *http.Request) {

	// Funcs servings as handlers take an http.ResponseWriter
	// and a http.REquest as args.
	// The response writer is used to fill in the HTTP response.
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// This handler reads all the HTTP request headers
	// and echos them into the response body
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// Register the handloers on server routes using
	// http.HandleFunc
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// call ListenAndServe with the port and handler
	// nil tells it to use the default router
	http.ListenAndServe(":8090", nil)
}