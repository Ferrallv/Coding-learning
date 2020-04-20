/*
Command line flags are a way to specify options for
command-line programs.
*/ 

package main


// The flag package supports basic command-line flag parsing
import (
	"flag"
	"fmt"
)

func main() {

	// Basic flag dec's can be used for string, int, and bool
	// options.
	wordPtr := flag.String("word", "foo", "a string")

	// Dec's numb and fork flags.
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// Dec an option that uses an existing var dec'd 
	// elsewhere in the program
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// After all flags are dec'd
	flag.Parse()

	// Dump out the parsed options
	// Need to dereference the pointers to get the 
	// actual values.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
