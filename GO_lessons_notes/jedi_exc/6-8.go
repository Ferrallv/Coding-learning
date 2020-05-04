package main

import "fmt"

func outside() func() {
	return func() {
		fmt.Println("This is the inside func")
	}
}

func main() {
	x := outside()

	x()
}
