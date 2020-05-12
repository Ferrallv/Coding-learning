// HANDS ON 2
// create a struct that holds person fields
// create a struct that holds secret agent fields and embeds person type
// attach a method to person: pSpeak
// attach a method to secret agent: saSpeak
// create a variable of type person
// create a variable of type secret agent
// print a field from person
// run pSpeak attached to the variable of type person
// print a field from secret agent
// run saSpeak attached to the variable of type secret agent
// run pSpeak attached to the variable of type secret agent
package main

import (
	"fmt"
)

type person struct {
	Name string
	Age int
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) pSpeak() {
	fmt.Println("What? My name is", p.Name)
}

func (s secretAgent) saSpeak() {
	fmt.Println("I can kill, my name is", s.Name)
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

	fmt.Println(reg_person.Age)
	reg_person.pSpeak()
	
	fmt.Println(agent.Name)
	agent.saSpeak()
	agent.pSpeak()
}