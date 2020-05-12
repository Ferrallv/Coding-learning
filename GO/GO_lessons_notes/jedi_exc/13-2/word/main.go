// Package word helps to analyze strings.
package word

import (
	"strings"
)

// UseCount returns a map of space separated words in a string.
func UseCount(s string) map[string]int {
	
	xs := strings.Fields(s)
	
	m := make(map[string]int)
	
	for _, v := range xs {
		m[v]++
	}
	
	return m
}

// Count returns the number of space seperated words in a string.
func Count(s string) int {
	count := 0
	x := strings.Split(s, " ")
	
	for range x {
		count++
	}
	return count
}