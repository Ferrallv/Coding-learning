package word

import (
	"fmt"
	"testing"
	"strings"
	"ExercisesAndProjectsInC/GO_lessons_notes/jedi_exc/13-2/quote"
)



func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
		xs := strings.Fields(quote.AriseArise)
		m := make(map[string]int)
	
		for _, v := range xs {
			m[v]++
		}
	}
}


func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
		count := 0
		x := strings.Split(quote.AriseArise, " ")
		
		for range x {
			count++
		}
	}
}


func ExampleUseCount () {
	for i, v := range UseCount("I I am fun fun fun") {
		fmt.Println(v, i)
	}
	// Output:
	// 2 I
	// 1 am
	// 3 fun
}

func ExampleCount() {
	fmt.Println(Count("I I am fun fun fun"))
	// Output:
	// 6
}

// Theres gotta be a better way! This would work
// if all v's were 0 and all i's didn't exist in 
// y!
func TestUseCount(t *testing.T) {
	x := UseCount(string_test)
	y := map[string]int{"I":2, "am":1, "fun":3}
	for i, v := range x {
		if v != y[i] {
			t.Error("got", v, "want", y[i])
		}
	}
}

func TestCount(t* testing.T) {
	x := Count(string_test)
	if x != 6 {
		t.Error("Got", x, "want", 6)
	}
}