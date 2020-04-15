/*
Range iterates ove elements in a variety of data 
structures
*/

package main

import "fmt"

func main() {
	// Let's use `range` to sum numbers in a slice.
	// An array will work the same way

	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)
}