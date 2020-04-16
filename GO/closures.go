/*
Go supports anonymous functions (PYTHON LAMBDA BIXIS)
They can form closures. A closure allows a function to access
captured variables (values or references to these values when
the closure was created) through the closure's copies, even when
the function is invoked outside of its scope. Whew!
*/

package main

import "fmt"

// The following function intSeq returns another function
// which is anonymously defined within intSeq. The returned
// function CLOSES OVER the variable i to form a 'closure'

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// We call `intSeq`, and assign the returned func to `nextInt`
// This func will capture its own `i` val and will update everytime
// we can nextInt. We can confirm the state is unique to each
// func by checking a new one!

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println("Check it out, it's a new func call!", newInts())
}