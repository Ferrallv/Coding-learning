package main

import "fmt"

func main() {
	m := map[string]int{
		"Nin" : 	1,
		"Piap" : 	2,
		"Duke" : 	3,
		"Count" : 	4,
	}
	fmt.Println(m)

	delete(m, "Nin")
	fmt.Println(m)

	// Can delete a key that doesn't exist
	// No error?
	delete(m, "Tyrande")
	fmt.Println(m)

	//  comma ok idiom
	if v, ok := m["Count"]; ok {
		fmt.Println("value:", v)
	}
}