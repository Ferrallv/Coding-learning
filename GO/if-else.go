/* Let's try some straightforward branching with 
if and else in GO!*/

package main

import ( 
	"fmt"
	"math/rand"
)

func main() {
	// A basic example to start
	var testing = rand.Intn(50)
	
	if testing % 2 == 0 {
		fmt.Println("The rando is even")
	} else {
		fmt.Println("The rando is odd")
	}

	// An `if` without and `else`, so lonely! But
	// comfortable in it's lonliness.
	if testing % 4 == 0 {
		fmt.Println("The rando is divisible by 4")
	}

	// We can have a statement precede conditionals and
	// any variable declared in this statement are available
	// in all branches.
	if num := rand.Intn(10) -2; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 3{
		fmt.Println(num, "is in range [0,2]")
	} else {
		fmt.Println(num, "is 4 or 5?")
	}
}