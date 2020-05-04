package main

import "fmt"

type person struct {
	first       string
	last        string
	favIceCream string
}

func main() {
	p1 := person{
		first:       "Aang",
		last:        "Avatar",
		favIceCream: "FrozenFrog",
	}

	p2 := person{
		first:       "Sokka",
		last:        "Water",
		favIceCream: "Mooncicle",
	}
	// slice1 := []string{p1}
	// slice2 := []string{p2}

	// for i := range slice1 {
	// 	fmt.Println(i)
	// }

	// for i := range slice2 {
	// 	fmt.Println(i)
	// }

	slice := []string{p1.favIceCream, p2.favIceCream}

	for _, v := range slice {
		fmt.Println(v)
	}
}
