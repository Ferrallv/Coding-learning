package main

import "fmt"

type puppy_paws int

var x puppy_paws

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}