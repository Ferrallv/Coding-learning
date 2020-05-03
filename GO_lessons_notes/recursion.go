package main

import "fmt"

func main() {
	n := factorial(4)
	fmt.Println(n)

	fmt.Println(forfact(4))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func forfact(n int) int {
	exclamation := 1
	for i := n; i > 0; i-- {
		exclamation = exclamation * i
	}
	return exclamation
}

//  4 * 3 * 2 * 1 * 0( STOP REPLACE WITH 1)