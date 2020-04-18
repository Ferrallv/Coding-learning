/*
A panic typically means something went unexpectedly wrong.
Is is used to fail fast on erros that shouldn't occur during
normal operations, or that we aren't prepared to handle gracefully
*/ 

package main

import "os"

func main() {
	// Using panic to check for unexpected errors.
	panic("a problem")

	// A common use of panic is to abort if a function returns an 
	// error vau
}