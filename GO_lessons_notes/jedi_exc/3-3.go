package main

import (
	"fmt"
	"time"
)

// Create a for loop with syntax `for condition {}`
// Have it print out the years you've been alive
func main() {
	birth := 1990
	for birth < time.Now().Year() {
		fmt.Printf("Vanin was alive in %d\n", birth)
		birth++
	}
}
