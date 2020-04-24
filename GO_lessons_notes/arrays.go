package main

import "fmt"

// Effective go states "Arrays are useful when planning the detailed layout of memory"
// And also "Use slices instead."

func main() {
	var x [5]int
	fmt.Println(x)
	
	x[3] = 23
	fmt.Println(x)
	fmt.Println(len(x))
}