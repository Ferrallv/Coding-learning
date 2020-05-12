package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 25; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case <-q:
			return
		case v := <-c:
			fmt.Println("This is from c:", v)
		}
	}
}
