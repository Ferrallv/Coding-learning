/*
Go supports pointers, allowing us to pass refreneces to values
and records within a program
*/

package main

import "fmt"

// Time to show (learn) pointers
// We will illustrate their difference to values
// zeroval will be a value. zeroptr will be our....pointer.
// zeroval will be passed an int value
// zeroptr will be pass an *iptr, which is an int pointer.

// I'll copy this for future reference

// the *iptr code in the function body then dereferences the
// pointer from its memory address to the current value at that 
// address. Assigning a value to a dereferenced pointer changes 
// the value at the referenced address

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}


func main() {
	i := 1
	fmt.Println("Initial val:", i)

	zeroval(i)
	fmt.Println("Zeroval doesn't change the i, it's out of 'scope'(?)", i)

	// The `&i` suntax gives the memory address of i, 
	// i.e. a pointer to i
	zeroptr(&i)
	fmt.Println("zeroptr changes the i, as it redirects the address to the assigned value, [should be 0]", i)

	// We can print the pointer as well
	fmt.Println("pointer", &i)
}