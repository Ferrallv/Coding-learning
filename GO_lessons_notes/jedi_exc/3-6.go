package main

import (
	"fmt"
	"time"
)

// Write a program to show the if statement in action
// 

func main() {
	if time.Now().Second() % 2 == 0 {
		fmt.Println("We are in an even second")
	} else {
		fmt.Println("We are in an odd second")
	}
}