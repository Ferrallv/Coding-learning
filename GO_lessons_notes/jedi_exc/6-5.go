package main

import (
	"fmt"
	"math"
)

type square struct {
	length float32
}

type circle struct {
	radius float32
}

func (s square) area() float32 {
	return s.length * s.length
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	// the func identifier and expected return are
	// both needed for the interface!
	area() float32
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	sqr := square{
		length: 10.2,
	}

	cir := circle{
		radius: 3.4,
	}

	info(sqr)
	info(cir)

}
