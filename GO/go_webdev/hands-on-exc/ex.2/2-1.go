package main

import (
	"fmt"
)

func main() {
	xi := []int{2,3,4,5,6}

	for i, _ := range xi {
		fmt.Println("Just the index: ", i)
	}

	for i, v := range xi {
		fmt.Println("Index and Value: ", i, v)
	}
}