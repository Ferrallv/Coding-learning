// for is Go's only looping construct. No while?
// okay!

package main

import "fmt"

func main () {
	// Start with the most basic for loop
	// It has a single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i+1
		// Look slike python so far!
	}

	// Classic initial/conditional/after `for` loop
	// j dec and inits to 7, loop while it's equal to or less than 9
	// increment j by 1 AFTER loop
	for j := 7; j <=9; j++ {
		fmt.Println(j)
	}

	// A `for` without a condition will loop repeatedly
	// until there is a `break` or a `return` from the enclosing
	// function.
	for {
		fmt.Println("loop")
		break //or it goes forever!
	}

	// `continue` is also available to skip to 
	// the next iter
	for n := 0; n <= 5; n++ {
		if n%2 == 0{
			continue
		}
		fmt.Println(n)
	} // This looks just like C and Python?
}

