package main

import "fmt"

// We don't say 'object' or 'class', but these
// are good approximations.
type person struct{
	first string
	last string
	age int
}

func main() {
	p1 := person{
		first: "James",
		last: "Bond",
		age: 32,
	}

	p2 := person{
		first: "Miss",
		last: "Moneypenny",
	}	

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2)
}