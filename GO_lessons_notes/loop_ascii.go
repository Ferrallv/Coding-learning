package main

import "fmt"

func main() {
	// 33-122
	for i := 33; i <= 122; i++{
		fmt.Printf("int: %d\tascii: %q\tunicode: %c\n", i, i, i)
	}
}