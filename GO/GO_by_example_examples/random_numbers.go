/*
pseudo-random number gen with math/rand
*/ 

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand.Intn returns int n, 0 <= n <  100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// rand.Float64 returns float64, f, 0.0 <= f < 1.0
	fmt.Println(rand.Float64())

	// Generate random floats in other ranges
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Default num gen is deterministc, so it'll produce
	// the same sequence of numbers each time by default.
	// To vary, give it a seed that changes. It is not safe
	// to use random numbers you intend to be secret, use
	// crypto/rand instead.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// If you seed a source with the same number, 
	// it produces the same sequence of rand nums
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
	fmt.Println()
}