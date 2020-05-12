/*
Parsing numnbers from strings
*/ 

package main

// Builtin package strconv provides num parsing
import (
	"fmt"
	"strconv"
)

func main() {
	// ParseFloat, 64 tells how many bits of precision to parse
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// ParseInt, 0 means infer base from string.
	// 64 requires that the result fit in 64 bits
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt will recognize hex-formatted nums
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// ParseUint
	u, _ := strconv.ParseUint("780", 0, 64)
	fmt.Println(u)

	// Atoi is for basic base-10 int parsing
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse funcs ret an error on bad input
	_, e :=strconv.Atoi("wat")
	fmt.Println(e)
}