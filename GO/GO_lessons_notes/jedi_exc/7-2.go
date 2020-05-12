package main

import "fmt"

// My solution, It worked! s
type person struct {
	height int
}

func changeMe(p *person) {
	p.height = 180
}

func main() {

	pers := person{
		height: 100,
	}
	fmt.Println(pers.height)
	fmt.Println(&pers)
	changeMe(&pers)
	fmt.Println(pers.height)
}

// this solution is provided by Todd.
// I ended up doing it correctly myself! Awesome!
// type person struct {
// 	name string
// }

// func main() {
// 	p1 := person {
// 		name: "James Bond",
// 	}
// 	fmt.Println(p1)
// 	changeMe(&p1)
// 	fmt.Println(p1)
// }

// func changeMe(p *person) {
// 	p.name = "Miss Moneypenny"
// 	// This works too.
// 	// (*p).name = "Miss Moneypenny"
// }
