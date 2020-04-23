 package main

 import "fmt"

 var y = 42
 
//  DECLARED the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the VALUE "shaken, not stirred"
var z string = "Shaken, not stirred"
var a string = `James said, 
"Shaken,

not stirred"`

// this is aSTATIC programming language
// a variable is DECLARED to hold a VALUE of a certain TYPE

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}