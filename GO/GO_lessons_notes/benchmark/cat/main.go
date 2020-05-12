package main

import (
	"fmt"
	"strings"
	"ExercisesAndProjectsInC/GO_lessons_notes/benchmark/cat/mystr"
)

const s = `On the twenty-fourth day of the month GORPIAIOS 2, which correspondeth to the twenty-fourth day of the fourth month of the season PERT 3 of the inhabitants of TA-MERT (EGYPT), in the twenty-third year of the reign of HORUS-RA the CHILD, who hath risen as King upon the throne of his father, the lord of the shrines of NEKHEBET 4 and UATCHET, 5 the mighty one of two-fold strength, the stablisher of the Two Lands, the beautifier of Egypt, whose heart is perfect (or benevolent) towards the gods, the HORUS of Gold, who maketh perfect the life of the hamentet beings, the lord of the thirty-year festivals like PTAḤ, the sovereign prince like RĀ, the King of the South and North, Neterui-merui-ȧtui-ȧuā-setep-en-Ptaḥ-usr-ka-Rā-ānkh-sekhem-Ȧmen 6, the Son of the Sun Ptolemy, the ever-living, the beloved of Ptaḥ, the god who maketh himself manifest. the son of PTOLEMY and ARSINOË, the Father-loving gods; when PTOLEMY, the son of PYRRHIDES, was priest of ALEXANDER, and of the Saviour-Gods, and of the Brother-loving Gods, and of the Beneficent Gods, and of the Father-loving Gods, and of the God who maketh himself manifest; when DEMETRIA, the daughter of Telemachus, was bearer of the prize of victory of BERENICE, the Beneficent Goddess; and when ARSINOË, the daughter of CADMUS, was the Basket Bearer of ARSINOË, the Brother-loving Goddess; when IRENE, the daughter of PTOLEMY, was the Priestess of ARSINOË, the Father-loving Goddess; on this day the superintendents of the temples, and the servants of the god, and those who are over the secret things of the god, and the libationers [who] go into the most holy place to array the gods in then apparel, and the scribes of the holy writings, and the sages of the Double House of Life, and the other libationers [who] had come from the sanctuaries of the South and the North to MEMPHIS, on the day of the festival, whereon His Majesty, the King of the South and North PTOLEMY, the ever-living, the beloved of Ptaḥ, the god who maketh himself manifest, the lord of beauties, received the sovereignty from his father,`

func main() {

	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))
}