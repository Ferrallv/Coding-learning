package main

import (
	"fmt"
)

func main() {
	the_ints := []int{}
	for i := 0; i < 11; i++ {
		the_ints = append(the_ints, i)
	}

	for _, v := range the_ints {
		if (v % 2) == 0 {
			fmt.Printf("%v is even\n", v)
		} else {
			fmt.Printf("%v is odd\n", v)
		}
	}
}