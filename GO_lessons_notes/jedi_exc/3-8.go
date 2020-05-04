package main

import "fmt"

// Create a program that uses a switch statemnet with no
// switch expression specified.

func main() {
	switch {
	case true:
		fmt.Println("This case is true and will print")
	case false:
		fmt.Println("This case is false and will not print")
	}
}
