/*
Go supports recursive functions. Let's go!
*/

package main

import "fmt"

// We create a `fact` func that calcs the factorial

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(8))
}