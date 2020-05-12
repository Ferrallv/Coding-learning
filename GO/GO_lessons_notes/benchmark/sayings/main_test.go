package sayings

import (
	"testing"
	"fmt"

)

func TestGreet(t *testing.T) {
	s := Greet("Nin") 
	if s != "Welcome my friend, Nin" {
		t.Error("got", s, "Expected `Welcome my friend, Nin`")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Nin"))
	// Output:
	// Welcome my friend, Nin
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Nin")
	}
}