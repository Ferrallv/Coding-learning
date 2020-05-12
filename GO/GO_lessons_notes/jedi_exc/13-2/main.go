package main

import (
	"fmt"
	"ExercisesAndProjectsInC/GO_lessons_notes/jedi_exc/13-2/word"
	"ExercisesAndProjectsInC/GO_lessons_notes/jedi_exc/13-2/quote"
)

func main() {
	fmt.Println(word.Count(quote.AriseArise))

	for k, v := range word.UseCount(quote.AriseArise) {
		fmt.Println(v, k)
	}
	fmt.Println(word.UseCount(quote.AriseArise) )
}