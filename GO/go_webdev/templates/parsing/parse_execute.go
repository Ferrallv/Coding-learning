package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates1/*"))
}

func main() {
	// tpl, err := template.ParseFiles("tpl.gohtml")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	/**************************/

	// tpl, err := template.ParseGlob("templates1/*")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	/**************************/

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}