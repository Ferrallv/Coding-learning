package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 23
}

func bar() (int, string) {
	return 20, "Is that all?"
}

// func receiever identifier paramater returns
