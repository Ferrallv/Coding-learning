/* `slices` are a data type in Go, a very powerful interface
to sequences, more so than arrays */

package main

import "fmt"

func main() {

	// Unlike arrays, slices are typed only by the elements they
	// contain (not the number of elements). To create an empty
	// slice with non-zero lenght, use the builtin `make`.
	// Starting to feel a lot like LISTS in Python!

	s := make([]string, 3)


	// We can "set" and "get" njust like arrays
	// Here I think "set" and "get" mean return the set, 
	// or an indexed return

	s[0] = "Hello, "
	s[1] = "World"
	s[2] = "!"
	fmt.Println("set :", s)
	fmt.Println("get :", s[2])

	// `len` returns the length of the slice
	fmt.Println("This is the `len`: ", len(s))

	// Slice supports a builtin `append`
	// append needs to accept a return value.
	s = append(s, "How")
	s= append(s, " are", " you?")
	fmt.Println("The append'd slice: ", s)


	// We can copy strings, but must first create
	// a string of the same length. Must it be Empty?
	c := make([]string, len(s))
	
	for i:=0; i < len(s); i++ {
		c[i] = "Strings?"
	}
	
	copy(c, s)
	fmt.Println("Trying the copy builtin :", c)

	// SLICES ARE PYTHON LISTS MOSTLY! 
	// They support list indexing, slicing a slice
	// like i[2:5]
	l := s[1:3]
	fmt.Println("Slicing the middle : ", l)

	l = s[:5]
	fmt.Println("Slice up to 5: ", l)

	l = s[3:]
	fmt.Println("Slice from 2 to the end: ", l)

	// Declare and initialize a slice in a single line
	t := []string{"goobs", "geebs", "gabs"}
	fmt.Println("Slice dec'd and init'd: ", t)

	// Like arrays, we can do multidimensional slices.
	twoDSlice := make([][]int, 3)
	for i:=0; i < 3; i++ {
		innerLen := i+1
		twoDSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDSlice[i][j] = i + j
		}
	}
	fmt.Println("Our 2D slice!", twoDSlice)
}