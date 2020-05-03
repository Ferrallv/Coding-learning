package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main() {
	
	// Anonymous struct created, as this struct doesn't
	// have a name! We declare the struct and define it all in
	// one go. 

	p1 := struct {
		first string
		last string
		age int
	}{
		first: "James",
		last: "Bond",
		age: 32,
	}
	

	// Trial? This does not work
	// p2 := struct {
	// 	first string
	// 	last string
	// 	age int
	// }
	// p2.first = "Aang"
	// p2.last = "Avatar"
	// p2.age = 112

	// Trial 2
	// This works, need to input empty interface. `{}`
	p2 := person{}

	p2.first = "Aang"
	p2.last = "Avatar"
	p2.age = 112
 

	fmt.Println(p1)
	fmt.Println(p2)
}