/*
context.Context for controlling cancellation.
A Context carries deadlines, cancellation signales, 
and other request-scoped vals accross API boundaries and
goroutines.
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// A context.Context is created for each request.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Wait before sending reply to client/
	// While working keep an eye on the context's Done()
	// channel to cancel work and return asap
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		// Context's Err() methodreturns an error that 
		// explains why the Done() channel was closed
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	// Register our handler on the "/hello" route
	// Start receiving
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}