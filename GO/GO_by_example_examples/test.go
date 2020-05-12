/*
Unit testing is an important part of writing principled Go
programs. The testing package provides the tools we 
need to write unit tests and the go test command runs tests

This code is in package main but it could be in any package
*/ 

package main

import (
	"fmt"
	"testing"
)

// Test the integer min finder func
func IntMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// A test is created by writing a func with a name
// beginning with Test
func TestIntMinBasic(t *testing.T) {
    ans := IntMin(2, -2)
    if ans != -2 {

 		// t.Error* will report failures but keep running
 		// t.Fail* will report test failures and stop the test
        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}

// Writing tests can be repetitive
// Use a table-driven style.
// Test inputs and expected outputs are listed in a table
// and a single loop walks over them and performs the test
// logic
func TestIntMinTableDriven(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

	// t.Run enables running "subtests", one for each table
	// entry. Shown seperately when executing go test -v

    for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
