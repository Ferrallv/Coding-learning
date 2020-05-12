/*
Go offers excellent support for string formatting in the
printf tradition.
*/ 

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var pf = fmt.Printf

func main() {

	// Go offers sever printing "verbs" designed to format
	// general Go values. For example, this prints an instance
	// of our point struct.
	p := point{1, 2}
	pf("%v\n", p)

	// If the val is a struct, the %+v variant includes
	// the struct's field names
	pf("%+v\n", p)

	// The %#v variant prints a Go syntax representation
	// of the value, ie the source code snippet that would
	// produce that val
	pf("%#v\n", p)

	// To print type of value, use %T
	pf("%T\n", p)

	// Formatting bools
	pf("%t\n", true)

	// base-10 int formatting
	pf("%d\n", 123)

	// Print out the binary representation
	pf("%b\n", 14)

	// For the char corresponding to this int
	pf("%c\n", 33)

	// %x for hex coding
	pf("%x\n", 456)

	// Basic decimal formatting for floats
	pf("%f\n", 78.9)

	// with different versions of scientific notation
	pf("%e\n", 123400000.0)
	pf("%E\n", 123400000.0)

	// For basic string printing
	pf("%s\n", "\"string\"")

	// To double qoute strings as in Go source
	pf("%q\n", "\"string\"")

	// using %x to render a string in base 16
	pf("%x\n", "hex this my homie!")

	// to print the representationof a pointer
	pf("%p\n", &p)

	// To specify the width and precision, of an int
	pf("|%6d|%6d|\n", 12, 345)

	// Of a float
	pf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// To left justify, use the - flag
	pf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// To align and control width when formatting strings
	pf("|%6s|%6s|\n", "foo", "b")

	// To left-justify
	pf("|%-6s|%-6s|\n", "foo", "b")

	// Printf prints the formatted string to os.Stdout
	// Sprintf formats and returns a string without pringting
	// it anywhere.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// you can format+print to io.Writers other than 
	// os.Stdout using Fprintf.
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}