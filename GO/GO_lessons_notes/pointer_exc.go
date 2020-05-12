package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x address before",&x)
	fmt.Println("x before", x)
	foo(x)
	bar(&x)
	fmt.Println("x address after", &x)
	fmt.Println("x address after", x)
}

func foo(y int) {
	fmt.Println(y)
	y = 23
	fmt.Println(y)
}

func bar(y *int) {
	fmt.Println(y)
	fmt.Println(*y)
	*y = 12
	fmt.Println(y)
	fmt.Println(*y)
}