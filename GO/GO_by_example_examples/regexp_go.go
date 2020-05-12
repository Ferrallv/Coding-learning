/*
Go offers builtin support for regular expressions.
*/ 

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var pl = fmt.Println

func main() {
	// Tests whether a pattern matches a string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	pl(match)

	// Some tasks require acomiled optimize Regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	pl(r.MatchString("peach"))

	// This finds the match for the regexp
	pl(r.FindString("peach punch"))

	// This also finds the first match but returns the start
	// and the end indexes for the match instead of the matching text
	pl(r.FindStringIndex("peach punch"))

	// The Submatch variant include info about both the 
	// whole-pattern and the subs within those matches.
	// This will return info for both p([a-z]+)ch and ([a-z]+)
	pl(r.FindStringSubmatch("peach punch"))

	// Indexes of matches and submatches
	pl(r.FindStringSubmatchIndex("peach punch"))

	// "All" variant funcs are available for the other
	// funcs as well
	pl(r.FindAllString("peach punch pinch", -1))

	pl(r.FindAllStringIndex("peach punch pinch", -1))

	// Providing a non-neg int as the second arg limits
	// the number of matches
	pl(r.FindAllString("peach punch pinch", 2))

	// You can provide []byte args and drop String from the 
	// func name
	pl(r.Match([]byte("peach")))

	// When creating global variables with regex you can use
	// the MustCompile variation. This panics instead of retunring
	// an error, which makes it safer to use for global vars.
	r = regexp.MustCompile("p([a-z]+)ch")
	pl(r)

	// regexp package can be used to replace subsets
	pl(r.ReplaceAllString("a peach", "<fruit>"))

	// The Func varian allows you to transform matched text
	// with a given function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	pl(string(out))
}	