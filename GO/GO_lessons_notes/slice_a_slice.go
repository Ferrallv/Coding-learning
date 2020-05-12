package main

import "fmt"

// Slicing a lot like in python.
func main() {
	x := []int{4, 5, 6, 7, 8, 23}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}