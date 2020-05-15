package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type meal struct{
	Name string
	Items []string
}

type menu []meal


func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := menu{
		meal{
			Name: "Breakfast", 
			Items: []string{"eggs",
				"bacon",
				"vegan tofu",
			},
		},
		meal{
			Name: "Lunch", 
			Items: []string{
				"sandwich",
				"soup",
				"ramYUM",
			},
		},
		meal{
			Name: "Dinner", 
			Items: []string{
				"steak",
				"cheese gnocchi",
				"dreams of Onerous Chupacabra",
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}