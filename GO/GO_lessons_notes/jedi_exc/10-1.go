package main

import (
	"fmt"
)

func main() {
	// Using func literal
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// Using buffered channel
	c2 := make(chan int, 1)

	c2 <- 234

	fmt.Println(<-c2)
}
