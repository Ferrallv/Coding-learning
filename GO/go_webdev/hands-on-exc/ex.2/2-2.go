package main

import (
	"fmt"
)

func main() {
	merp := map[string]int{
		"this": 2,
		"is": 4, 
		"cool": 6,
	}

	for k, _ := range merp {
		fmt.Println("Just the keys: ", k)
	}


	for k, v := range merp {
		fmt.Println("Keys and Vals ", k, v)
	}
}