package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base float64
}

func main() {
	tri := triangle{
		height: 12.5,
		base: 4.3,
	}

	squ := square{
		sideLength: 7.7,
	}

	fmt.Println(tri.getArea())
	fmt.Println(squ.getArea())

}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}