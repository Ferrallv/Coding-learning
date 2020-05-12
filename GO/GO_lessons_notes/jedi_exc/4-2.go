package main

import "fmt"

func main() {
	slice := []int{}

	for i := 0; i < 10; i++ {
		slice = append(slice, i+3)
	}

	for i := range slice {
		fmt.Println(slice[i])
	}
	fmt.Printf("%T\n", slice)
}
