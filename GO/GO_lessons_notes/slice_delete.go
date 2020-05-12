package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8, 23}
	fmt.Println(x)
	
	x = append(x, 77, 88, 99 , 1024)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}

	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}