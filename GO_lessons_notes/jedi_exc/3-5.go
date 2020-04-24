package main

import "fmt"

// Print out the modulus for each num between 10 and 100
// when divided by 4

func main() {
	for i := 10; i <= 100; i ++ {
		fmt.Printf("When %v is divided by 4, the modulus is %v\n", i, i%4)
	}
}