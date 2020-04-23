package main

import "fmt"

func main() {
	g := (2 == 4)
	h := (2 <= 4)
	i := (2 >= 4)
	j := (2 != 4)
	k := (2 < 4)
	l := (2 > 4)

	fmt.Println(g, h, i, j, k, l)
}