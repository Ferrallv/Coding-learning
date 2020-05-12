/*
Command line arguments are a common way to parameterize execution 
of programs.
*/ 

package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args provides access to raw command-line args
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// You can get individual args with indexing
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}