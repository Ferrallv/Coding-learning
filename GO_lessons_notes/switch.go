package main

import "fmt"

func main() {
	// switch {
	// case false:
	// 	fmt.Println("this should not print")
	// case (2 == 4):
	// 	fmt.Println("this should not print")
	// case (3 == 3):
	// 	fmt.Println("prints")
	// 	// fallthrough is a keyword to continue through to the 
	// 	// next case and run it regardless of statement result
	// 	fallthrough
	// case (4 == 4):
	// 	fmt.Println("Also true, does it print?")
	// 	fallthrough
	// case (7 == 9):
	// 	fmt.Println("not true")
	// 	fallthrough
	// case (1 == 4):
	// 	fmt.Println("not true 2")
	// 	fallthrough
	// case (15 == 15):
	// 	fmt.Println("True, for sure.")
	// 	fallthrough
	// default:
	// 	fmt.Println("This is default")
	// }

	switch "l" {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond james")
	case "Q", "M", "l":
		fmt.Println("this is Q, or M, or l.")
	default:
		fmt.Println("this is default")
	}
}