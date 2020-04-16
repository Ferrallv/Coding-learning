/*
Today I learned Vriadic Functions, a function with any 
number of trailing arguments. fmt.Println is a variadic function
*/

package main

import "fmt"

// let's write something that takes an arbitrary amount of `int`s

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Use variadic funcs the regular way we use funcs
	// with individual args
	sum(1, 2)
	sum(43, 6, 29)


	// If you have multi args in a slice you can 
	// apply them as args to a variadic func in the 
	// following way: `func(slice...)`
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
