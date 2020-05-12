package main

import "fmt"

// DECLARE the variable "y"
// ASSIGN the value 43
var y = 43

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
var z int

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x)
	
	// var will work outside of a func body
	// best practice is to keep scope as narrow as possible.
	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)
}