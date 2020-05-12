package main

import "fmt"

func main() {
	x := 10

	func(i int) {
		fmt.Println(i * i)
	}(x)
}
