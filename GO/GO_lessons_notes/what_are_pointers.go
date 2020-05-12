package main

import "fmt"

func main() {
	a := 23
	fmt.Println(a)
	// To see the address in memory
	fmt.Println(&a)
	
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	// Declaring b as a pointer to an int.
	// an address for memory. We looked up
	// the address of a using `&a` and wrote that
	// address to the variable b
	b := &a
	fmt.Println(b)
	// here we are dereferencing an address.
	// That is to say, what is in the address (the value)
	fmt.Println(*b)

	// The following two should have the same output?
	fmt.Println(*&a)
	fmt.Println(a)

	// assign the value saved at the address of b
	// (which is ALSO the address of a) to 42.
	*b = 42
	fmt.Println(a)
}

