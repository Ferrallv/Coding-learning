package main

import (
	"fmt"
)

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("8 + 12 =", mySum(8, 12))
	fmt.Println("3 + 7 =", mySum(3, 7))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
} 