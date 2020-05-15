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


type menus struct{
	 Restaurant string
	 Menu []meal
}


func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	rests := 
	menus{
		Restaurant: "Only One",
		Menu: []meal{
			{
				Name: "Breakfast", 
				Items: []string{"eggs",
					"bacon",
					"vegan tofu",
				},
			},
			{
				Name: "Lunch", 
				Items: []string{
					"sandwich",
					"soup",
					"ramYUM",
				},
			},
			{
				Name: "Dinner", 
				Items: []string{
					"steak",
					"cheese gnocchi",
					"dreams of Onerous Chupacabra",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, rests)
	if err != nil {
		log.Fatalln(err)
	}
}