package main

import "fmt"

func main() {
	x := []string{
		"James",
		"Bond",
		"Shaken, not stirred",
	}

	y := []string{
		"Miss",
		"Moneypenney",
		"Helloooooo, James",
	}

	multiD := [][]string{x, y}

	for i := range multiD {
		fmt.Println("Record Number:", i)
		for j := range multiD[i] {
			fmt.Printf("\t %s \n", multiD[i][j])
		}
	}
}