package main

import "fmt"

type vehicle struct {
	doors int
	colour string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {

	titan := truck{
		vehicle: vehicle{
			doors: 2,
			colour: "brown",
		},
		fourWheel: true,
	}

	camry := sedan{
		vehicle: vehicle{
			doors: 4,
			colour: "Champagne Sparkle",
		},
		luxury: false,
	}

	fmt.Println(titan, camry)
	fmt.Println(titan.doors)
	fmt.Println(camry.luxury)
}