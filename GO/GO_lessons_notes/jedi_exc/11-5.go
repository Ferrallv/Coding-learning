// Testing from Golang book
package main

import (
	"testing"
	"fmt"
)

func TestAverage(t *testing.T) {
	v := 9.0

	x := Average([]float64{1, 2, v})
	if x != 4.0 {
		t.Error("Expected 9, got ", x)
	}
}

func Average(q []float64) float64 {
	answer := 0.0
	for i := range q {
		answer += i
	}
	return answer
}