package main

import (
	"fmt"
)

func main() {
	s := "I'm sorry dave i can't do that"

	fmt.Println(s)
	fmt.Println([]byte(s))
	fmt.Println(string([]byte(s)))

	fmt.Println(s[:15])
	fmt.Println(s[10:23])
	fmt.Println(s[17:])
	for _, v := range s {
		fmt.Println(string(v))
	}
}