/*
A line filter is a type of program that reads input on stdin, 
processes it, then prints some derived result to stdout.
grep and sed are common line filters
*/ 

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Wrap unbuffered os.Stdin with a buffered scanner
	// This gives us a Scan method that advances the scanner
	// to the next token. This would be the next line in the 
	// default scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Text returns the current token.
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		// Write the uppercased line
		fmt.Println(ucl)
	}

	// Check for errors during scan. End of File
	// is expected and not reported by Scan as an error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}