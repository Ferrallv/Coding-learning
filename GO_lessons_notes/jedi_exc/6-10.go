package main

import "fmt"

func counter() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {

	a := counter()
	b := counter()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}