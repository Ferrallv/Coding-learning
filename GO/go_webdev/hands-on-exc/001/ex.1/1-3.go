// HANDS ON 3
// create an interface type that both person and secretAgent implement
// declare a func with a parameter of the interface’s type
// call that func in main and pass in a value of type person
// call that func in main and pass in a value of type secretAgent

package main

import (
	"fmt"
)

type person struct {
	Name string
	Age int
}

type secretAgent struct{
	person
	ltk bool
}

type human interface{
	Speak()
}

func (p person) Speak() {
	fmt.Println("What? My age is", p.Age)
}

func (s secretAgent) Speak() {
	fmt.Println("I can kill, my age is", s.Age)
}
 func info(h human) {
 	h.Speak()
 }


func main() {
	agent := secretAgent{
		person: person{
			Name: "Nin",
			Age: 24,
		},
		ltk: true,
	}

	reg_person := person{
		Name: "Aryn",
		Age: 43,
	}

	info(reg_person)
	info(agent)
}