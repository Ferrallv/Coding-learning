/*
In Go, there is support for multi return vals.
This feature is used often in idiomatic(?) Go, such as to 
return both the result and error falue from a function
*/

package main

import "fmt"


// the (intm int) signature shows that the func returns
// two `int`s
func vals() (int, int) {
	// Do lines matter in go? NOPE!
	return 3,
	



	7




}

func main() {

	//We can can use the 2 different return vals 
	// for multi assignment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// You can use the blank identifier if you only want
	// a subset of the returned vals
	c, _ := vals()
	fmt.Println(c)
}
