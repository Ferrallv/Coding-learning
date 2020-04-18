// Go's structs are typed collections of fields
// It is useful to group data together to form records.
// What?

package main

import "fmt"

// Creating a `person` struct, with `name` and `age` fields.

type person struct {
	name string
	age int
}

// newPerson constructs a new `person` struct with
// the given name.

// You can safely return a pointer to a local variable
// as a local variable will survive the scope of the function

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// Create a new struct
	fmt.Println(person{"Alpha", 29})

	// You can name the fields when init'ing a struct
	fmt.Println(person{name: "Beta", age: 27})

	// Omitted fields are zero-valued
	fmt.Println(person{name: "Gamma"})

	// an & will yield a pointer to the struct
	fmt.Println(&person{name:"Omega", age: 21})

	// It's idiomatic to encapsulate new struct creation
	// in constructsor functions.
	fmt.Println(newPerson("Epsilon"))

	// Access struct fields with a dot (like python classes?)
	s := person{name:"Zeta", age: 50}
	fmt.Println(s.name)

	// You can use dots with struct pointers
	// the pointers are automatically dereferenced.
	// ** Dereferencing means getting the value that is 
	// stored in the memory location that is pointed by
	// the pointer

	sp := &s
	fmt.Println(sp.age)

	// Structs are MUTABLE
	fmt.Println("Original", sp.name, sp.age)
	sp.name, sp.age = "Thomas", 21

	fmt.Println("After", sp.name, sp.age)

}