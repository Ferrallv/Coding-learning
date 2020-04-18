/* Switch statement: conditionals 
across many branches */

package main

import (
	"fmt"
	"time"
)

func main() {
	// A basic switch
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// You can use commas to separate multiple expressions
	//in the same case statement. There us also the `default`
	//case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend") // This is so cool!
	default: 
		fmt.Println("It's a weekday")
	}

	// `switch` without an expression is an alternate way to 
	//express if/else logic. Here we also show how we can have
	// a switch expression can be a non-constant. Wow!
	t := time.Now()
	switch {
	case t.Hour() < 17 || t.Hour() > 9: //just like python!
		fmt.Println("It's still during work hours")
	default:
		fmt.Println("We be closed.")
	}

	/*We can use a type switch! Us this to discover
	the type of interface value (?) */
	whatAmI := func(i interface{}) {
		switch t:= i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	// We created a function! Let's use it!
	whatAmI(false)
	whatAmI(32)
	whatAmI("Is this a string?")

}