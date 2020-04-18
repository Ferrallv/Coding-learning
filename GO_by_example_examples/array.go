/*ARRAYS BOO YA*/

package main

import "fmt"

func main() {
	// Creating an array that holds EXACTLY 5 ints.
	// Unlike python, I htink an array in go cannot change
	// len or type.
	//default array is zero-valued, for ints that is 0 
	var a [5]int
	fmt.Println("An empty array: ", a)

	// Set the value at in index of an array 
	//just like in python `arra[index] = blah`
	a[4] = 31
	fmt.Println("the whole set:", a)
	fmt.Println("just get indexed value, here it's a[4]", a[4])

	// LIKE IN PYTHON WE CAN GET THE LEN WITH A BUILTIN
	fmt.Println("len:", len(a))

	// This syntax is used to declare and initialize
	// an array in one line.
	b := [5]int{10, 9, 8 , 7, 6}
	fmt.Println("Our dec'd and init'd array: ", b)

	// Array types are one-dimensional, but you can compose
	// types to build multi-dimensional data structures.
	var twoDarray [2][3]int
	for i := 0; i<2; i++{
		for j :=0; j<3; j++ {
			twoDarray[i][j] = i+j
		}
	}
	fmt.Println("That super cool 2D array we made :", twoDarray)
}

