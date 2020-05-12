package main

import "fmt"

const a int = 23

const b = "A String?"

// OR

const (
	c        = true
	d string = "It is"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(c, d)
}
