package main

import (
	"fmt"
	"time"
)

// Build on 3-6 to use else if
// 

func main() {
	if time.Now().Minute() % 5 == 0 {
		fmt.Println("We are starting a new twelfth of an hour")
	} else if time.Now().Minute() % 2 == 0  {
		fmt.Println("We are in an even minute")
	} else {
		fmt.Println("We are in an odd minute")
	}
}