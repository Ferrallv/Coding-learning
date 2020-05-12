package main

import "fmt"

func main() {
	ii := []int{1,2,3,4,5,6,7,8,9,}
	s := sum(ii...)
	fmt.Println("sum of all numbers", s)

	s2 := evens(sum, ii...)
	fmt.Println("sum of all even numbers", s2)

	s3 := odds(sum, ii...)
	fmt.Println("sum of all even numbers", s3)
}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//  This is a callback
func evens(f func(xi ...int) int, vi ...int) int {
	var yi []int 
	for _, v := range vi {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odds(f func(xi ...int) int, vi ...int) int {
	var zi []int
	for _, v := range vi {
		if (v % 2) != 0 {
			zi = append(zi, v)
		}
	}
	return f(zi...)
}