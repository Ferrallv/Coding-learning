// Package dog allows us to more fully understand dogs.
package dog

import (
	"testing"
)

const the_year int = 5

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++{
		Years(the_year)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++{
		YearsTwo(the_year)
	}
}

func ExampleYears() {
	Years(5)
	// Output:
	// 30
}

func ExampleYearsTwo() {
	Years(10)
	// Output:
	// 70
}

func TestYears(t *testing.T) {
	y := Years(the_year)
	if y != 35 {
		t.Error("Got", y, "wanted 35")
	}
}

func TestYearsTwo(t *testing.T) {
	y := YearsTwo(the_year)
	if y != 35 {
		t.Error("Got", y, "wanted 35")
	}
}