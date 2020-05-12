/*
Sometimes we'll want to sort a collection by something other
than its natural order. What if we wanted to sort strings
by their len?
*/ 

package main

import (
	"fmt"
	"sort"
)

// In order to sort by a custom function in Go, we need
// a corresponding type. Here we've created a byLength type
// that is just an alias for the builtin []string type
type byLength []string

// We implement sort.Interface - Len, Less, and Swap - on
// our type so we can use the sort package's generic Sort
// func. Len and Swap will usually be similar across types
// and Less will hold the actual custom sorting logic.
// In this case we want to sort in order of increasing
// string len.
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// With this in place, we can now implement our custom
// sort by converting the original fruits slice to byLength
// and then use sort.Sort on that typed slice.
func main() {
	fruits := []string{"banana", "apricot", "apple", "pear", "shitburgers", "fig"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}