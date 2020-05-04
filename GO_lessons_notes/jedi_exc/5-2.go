package main

import "fmt"

type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {
	p1 := person{
		first: "Aang",
		last:  "Avatar",
		favIceCream: []string{
			"FrozenFrog",
			"Veggie",
			"AirWhipped",
		},
	}

	p2 := person{
		first: "Sokka",
		last:  "Water",
		favIceCream: []string{
			"Mooncicle",
			"Boomerang",
			"Sooki",
		},
	}

	// for _, v := range p1.favIceCream {
	// 	fmt.Println(v)
	// }

	// for _, v := range p2.favIceCream {
	// 	fmt.Println(v)
	// }

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Println(k)
		for i, v := range v.favIceCream {
			fmt.Println(i, v)
		}
	}

}
