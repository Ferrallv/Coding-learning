/*
Functions are central in Go. Is this untrue in 
other languages? 
*/

package main

import "fmt"

// A function that takes two `int`s and returns
// their sum as an `int`

func plus(a int, b int) int {

	// Go requires explicit returns, it won't 
	// automatically return the value of the last expression
	return a + b
}

/*
Below we can see that if you have consecutive like typed
parameters, you can declare the type on the last parameter
for all of them
*/

func plusPlus(a, b, c int) int {
	return a + b + c
}

// We can call functions as expected with `name(args)`

func main() {

	res := plus(10, 12)
	fmt.Println("10 + 12 =", res)

	res = plusPlus(10, 9 , 8)
	fmt.Println("what is the plusPlus?", res)
}