package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	
	type test struct {
		data []int
		answer int
	}

	tests := []test{
		test{[]int{10,10,3}, 23},
		test{[]int{1,2,3}, 6},
		test{[]int{10,10}, 20},
		test{[]int{7,7,7}, 21},
	}

	for _, v := range tests {
		x := mySum(v.data...) 
		if x != v.answer {
			t.Error("Expected",v.answer,"got", x)
		}
	}


}