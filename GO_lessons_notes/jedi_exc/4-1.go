package main

import "fmt"

func main() {
	arr := [5]int{}

	for i := 0; i < len(arr); i++ {
		arr[i] = 23
	}

	for i := range arr {
		fmt.Println(arr[i])
	}

	fmt.Printf("%T\n ", arr)
}