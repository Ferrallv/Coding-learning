package main

import (
	"fmt"
)

type customErr struct {
	ourError string
}

func (ce customErr) Error() string {
	return ce.ourError
}

func main() {
	e := customErr{
		ourError: "Is this a center for ants?",
	}
	foo(e)
}

func foo(e error) {
	fmt.Println(e)
} 