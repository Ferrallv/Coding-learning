package mymath

import (
	"testing"
	"fmt"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++{
		CenteredAvg([]int{1,2,2,3,3,4})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1,2,4,6,9}))
	// Output:
	// 4
}
  
func TestCenteredAvg(t *testing.T) {
	type test struct {
		data []int
		answer float64
	}

	tests := []test{
		test{[]int{1,2,3,4,5}, 3},
		test{[]int{1,2,2,2,15}, 2},
		test{[]int{-4 ,4,4,4,236}, 4},
		test{[]int{7,7,7}, 7},
		test{[]int{123,123,123,4}, 123},
	}


	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.answer {
			t.Error("Got", x, "want", v.answer)
		}
	}
}
