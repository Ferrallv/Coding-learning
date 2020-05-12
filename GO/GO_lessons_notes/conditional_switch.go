package main

import "fmt"

func main() {
	// It is idiomatic to write an if-else-if-else chain
	// as a switch
	switch {
	default:
		fmt.Println("testing default at the start")
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}