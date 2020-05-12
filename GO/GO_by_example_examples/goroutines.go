// a goroutine is a lightweight thread of execution(?)

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Suppose we have a function call f(s). Here's how we'd
	// call that in the usual way, running synchronously
	f("direct")

	// To invoke this function in a goroutine, use go `f(s)`
	// This new goroutine will execute concurrently with the 
	// calling one.
	go f("goroutine")

	// You can also start a goroutine for an anonymous 
	// function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously
	// in spearate goroutines now. We wait for them to finish

	time.Sleep(time.Second)
	fmt.Println("done")
}
