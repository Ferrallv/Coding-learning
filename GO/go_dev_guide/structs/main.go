package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	pCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}

	// sonya := person{firstName: "Sonya", lastName: "Simpson"}

	// fmt.Println(alex, sonya)
	// fmt.Printf("First name: %v, Last Name: %v", alex.firstName, alex.lastName)

	// fmt.Printf("Your name? %+v", alex)

	jim := person{
		firstName: "Nin",
		lastName:  "Ja",
		contact: contactInfo{
			email: "nin@help.com",
			pCode: "N1N4E1",
		},
	}
	
	jimPointer := &jim
	jimPointer.updateName("Sally")
	jim.Print()
}

func (p *person) updateName(newFirstName string) {
	  p.firstName = newFirstName

}

func (p person) Print() {
	fmt.Printf("%+v\n", p)
}
