package main

import "fmt"

func main() {
	xi := []int{1,2,3,4,5,6,7,8,9}
	
	// Those dots are the unfurl/unpack.
	x := sum(xi...)
	fmt.Println("The total is:", x)
}

// func (r receiver) identifier(paramter(s)) (return(s))
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0 
	for i, v := range x {
		sum += v
		fmt.Println("For item in index position,", i, "we are now adding, ", v, " To the total: ", sum)
	}
	fmt.Println("The total is: ", sum)
	return sum
}