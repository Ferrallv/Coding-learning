package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "- secretAgent Speak")
}

type human interface {
	speak() 
}


// Assertion
func bar(h human) {
	switch h.(type){
	case person:
		fmt.Println("I was passed into barrrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - Person Speak")
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}
	
	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last: "Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Doctor",
		last: "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}