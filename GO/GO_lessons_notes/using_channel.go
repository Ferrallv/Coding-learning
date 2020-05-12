package main

import (
	"fmt"
)

func main() {
	
	c := make(chan int)
	
	// send
	go foo(c)
	
	// receive
	bar(c)

	// A note:
	/*
	bar is not used a s a go routine. If both are off
	running in parallel, the main function might finish
	before either is complete. By making the bar chan block
	we are able to wait for bar() before main() finishes.
	*/ 
	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
