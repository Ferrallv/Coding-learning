package main

import (
	"io"
	"net/http"
	"html/template"
	"log"
)

var tpl *template.Template

type person struct {
	Name string
	Age int
}

func landing(w http.ResponseWriter, req *http.Request) {
	nin := person{
		Name: "Nin Ja",
		Age: 23,
	}
	err := tpl.ExecuteTemplate(w, "index.gohtml", nin)
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "My dog is Bennett")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am Nin")
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", landing)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}