package main

import (
	"fmt"
	"math"
)


// const declares a constant value. It has global scope?
const dec_con_value = "constant"

func main() {
	fmt.Println(dec_con_value)

	// We can drop const where ever we can drop var
	const a_num = 50000000

	// Constants permform arithmetic with arbitrary precision.
	// Interesting
	const calc_it_up = 3e20 / a_num
	fmt.Println(calc_it_up)

	// Numeric const has no type till it's given,
	// such as through explicit conversion
	fmt.Println(int64(calc_it_up))

	/* A number can be given a type by using it
	in a context where it is needed to have a specific 
	type. Here `math.Sin` expects a float, and she will
	get a float! */
	fmt.Println(math.Sin(calc_it_up))
}