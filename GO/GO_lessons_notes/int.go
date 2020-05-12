package main

import "fmt"

var x int

var y float64

var x1 int8 = -128


func main() {
	x = 42
	y = 42.3121

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(x1)
	fmt.Printf("%T\n", x1)
}