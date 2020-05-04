package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// func (r receiver) identifier(parameters) (return(s)) { code }

func (p person) Speak() {
	fmt.Println("I am", p.first, " and I am", p.age, " years old.")
}

func main() {

	person1 := person{
		first: "Nin",
		last:  "Aeve",
		age:   34,
	}

	person1.Speak()
}
