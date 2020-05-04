package main

import "fmt"

// Use a switch statement with the switch expression specified
// as a variable of TYPE string with the identifier "favSport"

func main() {
	favSport := "cycling"

	switch favSport {
	case "monkey":
		fmt.Println("Monkey is not a sport")
	case "cycling":
		fmt.Println("nice")
	}
}
