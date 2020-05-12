package main

import "fmt"

func main() {
	x := []int{}

	for i := 42; i < 52; i++ {
		x = append(x, i)
	}

	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
