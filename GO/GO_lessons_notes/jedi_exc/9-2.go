package main

import (
	"fmt"
)

type person struct {
	Name string
}

func (p *person) Speak() {
	fmt.Println(p.Name)
}

type human interface {
	Speak()
}

func saySomething(h human) {
	h.Speak()
}

func main() {
	myPerson := person{
		Name: "Nin",
	}

	// This will work
	saySomething(&myPerson)

	// This will no work
	// saySomething(myPerson)
}
