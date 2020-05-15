package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
	favFood []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type transportation interface {
	transportationDevice() string
}

func (p person) walk() string {
	return fmt.Sprintf("%v is walking", p.fName)
}

func (s sedan)transportationDevice() string {
	return fmt.Sprintf("I carry families, not much else")
}

func (t truck)transportationDevice() string {
	return fmt.Sprintf("I carry everything")
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func main() {
	p1 := person{
		fName: "Rand",
		lName: "TwoRivers",
		favFood: []string{"wool", "trollocs", "AesSedai"},
	}

	fmt.Println(p1.favFood)
	for _, v := range p1.favFood {
		fmt.Println(v)
	}

	s := p1.walk()
	fmt.Println(s)

	fonefifty := truck {
		vehicle : vehicle {
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	camry := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "champagne sparkle",
		},
		luxury: false,
	}

	fmt.Println(fonefifty)

	fmt.Println(camry)
	fmt.Println(camry.doors)
	fmt.Println(camry.color)
	fmt.Println(camry.luxury)

	report(camry)
	report(fonefifty)

}