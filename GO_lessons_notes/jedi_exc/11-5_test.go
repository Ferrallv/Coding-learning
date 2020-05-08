// Testing from Golang book
package main

import (
	"testing"
)

func TestAverage(t *testing.T) {
	v := 6.0

	x := Average([]float64{1.0, 2.0, v})
	if x != 3.0 {
		t.Error("Expected 3.0, got ", x)
	}
}

func Average(q []float64) float64 {
	total := 0.0
	for i := range q {
		total += float64(i)
	}
	return total
}

// Something isn't right....