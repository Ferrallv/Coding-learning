package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(slice[:5])
	fmt.Println(slice[5:])
	fmt.Println(slice[2:7])
	fmt.Println(slice[1:6])
}
