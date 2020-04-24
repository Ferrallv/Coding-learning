package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8, 23}
	fmt.Println(x)
	
	x = append(x, 77, 88, 99 , 1024)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	// The triple dots opens up the slice and appends the values
	// within.
	z := append(x, y...)
	fmt.Println(z)

	// Try without the dots.
	// x = append(x, y)
	// fmt.Println(x)
	// IT DOES NOT WORK! You need to unpack the slice.
}