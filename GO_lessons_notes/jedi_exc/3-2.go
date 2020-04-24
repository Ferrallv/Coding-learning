package main

import "fmt"

// Print every rune code of the uppercase alphabet three times

func main() {
	// lol 65+25 is 90 you knob.
	for i := 65; i <=(65+25); i++{
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			// fmt.Printf("\t%U %q\n", i, i)
			fmt.Printf("\t%#U\n", i)
			}
	}
}