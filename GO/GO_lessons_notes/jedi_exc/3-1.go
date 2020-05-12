package main

import "fmt"

// print every number from 1 - 10,000

func main() {
	for i := 1; i <= 10000; i++ {
		fmt.Printf("%v\t", i)
	}
}
