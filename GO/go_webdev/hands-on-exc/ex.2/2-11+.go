package main

import (
	"fmt"
)

type gator int

type flamingo bool

type swampCreature interface {
	greeting()
}

func (g gator) greeting() {
	fmt.Println("Hello, I am gator")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")
}

func bayou(s swampCreature) {
	s.greeting()
}


func main() {
	var g1 gator
	var x int
	var f1 flamingo
	
	f1 = true
	g1 = 23
	
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// You CAN assign x to g1, if they share the same
	// type. You can convert the type, if the underlying type
	// is the same
	g1 = gator(x)

	g1.greeting()
	f1.greeting()

	bayou(g1)
	bayou(f1)

}