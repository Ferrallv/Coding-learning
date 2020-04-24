package main

import "fmt"

// maps are key value store, much like python dict
func main() {
	// the COMPOSITE LITERAL is `map[string]int{}`
	// This is the TYPE.
	m := map[string]int{
		"James": 			32,
		"Miss Moneypenny": 	27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	// Will return 0 if there is no entry
	fmt.Println(m["Barnabas"])

	// The comma ok idion
	v, ok := m["Baranabas"]
	fmt.Println(v)
	fmt.Println(ok)

	

	if v, ok := m["James"]; ok {
		fmt.Println("THIS IS THE PRINT", v)
	}
}