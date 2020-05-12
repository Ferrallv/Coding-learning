package main

import "fmt"

func main() {

	// ints1 := []int{1,2,3,4,5}
	fmt.Println(foo([]int{1, 2, 3, 4, 5}...))
	fmt.Println(bar([]int{2, 4, 6, 8}))

}

func foo(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func bar(ii []int) int {
	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum

}
