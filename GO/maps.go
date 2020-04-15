/*
Maps are Go's built-in associative data-type. Like Python's
dict! Key-value pairs, here we come!
*/

package main

import "fmt"

func main() {

	// Creating an empty map is easy!
	// Follow the syntax for the builtin:
	//	make(map[key-type]val-type).
	m := make(map[string]int)

	// Set the key:value pairs using same as python
	//syntax. SuhWheet!
	m["Code"] = 1337
	m["Agent"] = 7

	//checkit
	fmt.Println("This is our map: ", m)

	// Get a value for a key with `name[key]`

	v1 := m["Code"]
	fmt.Println("Our v1 var we dec'd and init'd: ", v1)

	// Builtin `len` returns the number of pairs in our
	// map
	fmt.Printl n("This is the number of key/value pairs in our map: ", len(m))

	// Builtin `delete` removes key/value pairs from the map
	delete(m, "Agent")
	fmt.Println("What is left in our map? ", m)

	/*
	When we add the optional second return value when we get
	a value form a map it is used to indicate if the key was
	present in the map. Can be used to 'disambiguate' between missing
	keys and keys with 0 or blank values (""). IF you don't want
	the value itself you can use an ignore option with the 
	blank identifier `_`
	*/
	_, pairs := m["Agent"]
	fmt.Println("Is thar a par of pairs? ", pairs)

	// To dec and init a new map in the same line
	n := map[string]int{"Poo":1, "pee":2}
	fmt.Println("The New Map Kids: ", n)
}