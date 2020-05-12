package main

import "fmt"

func main() {
	foo()
	bar("Nin")
	s1 := woo("Piap")
	fmt.Println(s1)
	x, y := mouse("Duke", "Count")
	fmt.Println(x)
	fmt.Println(y)
}


// Creating a function
// func (r receiver) identifier(parameters) (return(s)) {...}
// Define a func with parameters(if any).
// Call a func and pass in args (if any).

func foo() {
	fmt.Println("hello from foo")
}

// Everything in Go is PASS BY VALUE - it's WYSIWYG
func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo ", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " & ", ln, ` say, "Hello"`)
	b := false
	return a, b
}