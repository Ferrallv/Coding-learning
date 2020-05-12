package main

import "fmt"

func main() {
	// create and use an anon struct

	nin := struct {
		height     float32
		weight     float32
		sunglasses bool
		likes      map[string]bool
		skills     []string
	}{
		height:     1.87,
		weight:     92.0,
		sunglasses: true,
		likes: map[string]bool{
			"Pop":          false,
			"Candy":        false,
			"PeanutButter": true,
		},
		skills: []string{
			"Lifting",
			"Sighing",
			"Cuddling Dogs",
		},
	}

	fmt.Println(nin.height, nin.weight)
	fmt.Println(nin.sunglasses)

	for k, v := range nin.likes {
		fmt.Println(k, v)
	}

	for _, v := range nin.skills {
		fmt.Println(v)
	}
}
