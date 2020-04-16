/*
Range iterates ove elements in a variety of data 
structures
*/

package main

import "fmt"

func main() {
	// Let's use `range` to sum numbers in a slice.
	// An array will work the same way

	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	/*
	`range` on arrays and slices also provides the index.
	In the last loop we ignored the index with the blank 
	identifier. Sometimes we want the index though
	*/

	for i, num := range nums {
		if num == 3 {
			fmt.Println("This is the index:", i)
		}
	}

	// `range` on a map will iterate over the key/value
	// pairs

	key_vals := map[string]string{"a": "apple", "b": "banana"}
	for key, val := range key_vals {
		fmt.Printf("%s -> %s\n", key, val)
	}

	// `range` can iterate over just the keys too
	for key := range key_vals {
		fmt.Println("The Key:", key)
	}

	// `range` can iterate over a string as well! It iterates
	// over the Unicode code points. THe first value is the starting
	// byte index of the `rune`	and the second is the `rune` itself.
	// uh, Cool!
	 for index, rune := range "go" {
	 	fmt.Println(index, rune)
	 }
}