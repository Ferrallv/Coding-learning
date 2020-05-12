package main

import "fmt"

func main() {
	m := map[string]int{
		"Nin": 		1,
		"Piap":		2,
	}

	m["Duke"] = 3
	m["Count"] = 4

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{1,2, 3, 4}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}