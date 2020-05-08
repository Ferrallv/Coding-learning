package word

import (
	"fmt"
	"testing"
	"ExercisesAndProjectsInC/GO_lessons_notes/jedi_exc/13-2/quote"
)



func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.AriseArise)
	}
}


func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.AriseArise)
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
	x := UseCount("I I am fun fun fun")
	y := map[string]int{"I":2, "am":1, "fun":3}
	for i, v := range x {
		if v != y[i] {
			t.Error("got", v, "want", y[i])
		}
	}
}

func TestCount(t* testing.T) {
	x := Count("I I am fun fun fun")
	if x != 6 {
		t.Error("Got", x, "want", 6)
	}
}