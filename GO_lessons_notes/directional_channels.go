package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 23
	}()

	fmt.Println(<-c)
	fmt.Println("-------")
	
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send


	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// General to specific converts
	fmt.Println("-------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	// Specific to general doesn't convert
	// fmt.Println("-------")
	// fmt.Printf("c\t%T\n", (<-chan int)(cr))
	// fmt.Printf("c\t%T\n", (chan<- int)(cs))
}