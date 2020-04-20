/*
How to set, get, and list env vars
*/ 

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// To set a key/value pair, use os.Setenv.
	// To get a val for a key, use os.Getenv.
	// It will return an empty string if the key isn't 
	// present in the env.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR", os.Getenv("BAR"))

	// Use os.Environ to list all key/value pairs in the env
	// This returns a slice of strings in the form KEY=val.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
