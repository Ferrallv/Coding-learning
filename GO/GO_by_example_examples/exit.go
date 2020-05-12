/*
os.Exit to immediately exit with a given status
*/ 

package main

import (
	"fmt"
	"os"
)

func main() {
	// defers will not run when using os.Exit
	// The fmt will never be called
	defer fmt.Println("!")

	// Exit with status 3
	os.Exit(3)
}
e