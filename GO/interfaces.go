// Interfaces are named collections of method signatures
// Is an interface a class? Yes and no!

package main

import (
	"fmt"
	"math"
)

// A basic interface for geometric shapes
type geometry interface {
	area() float64
	perim() float64
}

// 