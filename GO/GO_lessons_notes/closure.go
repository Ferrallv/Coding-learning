package main

import "fmt"



func main() {
	// WOW! This func "holds" or encloses a value
	// we can call it multiple times to increment a value
	a := incrementer()
	b := incrementer()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementer() func() int {
	var x int
	return func() int {
			x++
			return x
	}
}