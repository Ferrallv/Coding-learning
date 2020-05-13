package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Func(fm).ParseFile("tpl.gohtml"))
}

fun main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
}