package main

import "fmt"

func main() {
	dict := map[string][]string{
		`bond_james`:       []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenney_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:            []string{`Being Evil`, `Ice cream`, `Sunsets`},
	}

	// I originally wrote what is below, but agree that above is more readable!
	// dict["bond_james"] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	// dict["moneypenney_miss"] = []string{`James Bond`, `Literature`, `Computer Science`}
	// dict["no_dr"] = []string{`Being Evil`, `Ice cream`, `Sunsets`}

	for k, v := range dict {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Printf("\t %v %v\n", i, v2)
		}
	}
}
