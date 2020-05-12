package main

import "fmt"

func main() {
	// init; condition; post;
	for i := 0; i <= 20; i++ {
		fmt.Println(i)
	}

	// Nested loop
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++{
			 fmt.Printf("\t\tThe inner loop: %d\n", j)
		}
	}

}