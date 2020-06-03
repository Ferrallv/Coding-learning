package main

import (
	"fmt"
)

type bot interface {
	GetGreeting() string
}

type englishBot struct {}

type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	PrintGreeting(eb)
	PrintGreeting(sb)
}

func PrintGreeting(b bot) {
	fmt.Println(b.GetGreeting())
}

// We can omit the value of the receiver if we are 
// not using it in our function
func (englishBot) GetGreeting() string {
	return "Hi There!"
}

func (spanishBot) GetGreeting() string {
	return "Hola!"
}