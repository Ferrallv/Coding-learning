package main

import "fmt"

func main() {

	var it_a_string = "initial, This is automatically a string. Lines don't matter in Go? If this works they do"
	fmt.Println(it_a_string)

	// Declareing multiple variables, nice!
	// Will int force a float to int here too?
	// Compiler notice that 12.2 is truncated, and does not run
	// with go run variables.go
	var b, c int = 300, 12  //12.2
	fmt.Println(b, c)

	// Go infers variable types, awesome!
	var yes_is_true = true
	fmt.Println(yes_is_true)

	// This shiz is zero valued if we declare but
	// initialize. Fancy
	var e int
	fmt.Println(e)

	// Demonstratitng := shorthad for declaring
	// and intializing. This is spice.
	hats := "tricks"
	fmt.Println(hats)
}