package main

import (
	"fmt"
	"time"
)

// Create a for loop using syntax `for {}`
func main() {
	bd := 1990
	for {
		if bd <= time.Now().Year() {
			fmt.Println(bd)
			bd++
		} else {
			fmt.Println("Done")
			break
		}
	}
}
