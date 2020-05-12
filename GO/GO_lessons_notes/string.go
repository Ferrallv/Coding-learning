package main

import (
	"fmt"
)

 func main() {
 	s := "Hello, playground"
 	fmt.Println(s)
 	fmt.Printf("%T\n", s)

 	bs := []byte(s)
 	fmt.Println(bs)
 	fmt.Printf("%T\n", bs)

 	for i := 0; i < len(s); i++ {
 		fmt.Printf("%#U ", s[i])
 	}

 	fmt.Println("")

 	for i, v := range s {
 		fmt.Printf("at index position %d we have hex %#x\n", i, v)
 	} 

 	// hex
 	s1 := "H"
 	fmt.Println(s1)

 	bs1 := []byte(s1)
 	fmt.Println(bs1)
 	
 	n := bs1[0]
 	fmt.Printf("%T\n", n)
 	fmt.Printf("%b\n", n)
 	fmt.Printf("%#x\n", n)
 }