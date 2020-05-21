package main

import (
	"io"
	"net/http"
	"html/template"
	"log"
)


func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/Golde33443.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	
}

// Notice we need a separate function to `serve` the 
// image, then main grabs it and sends it to the template html.

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "Golde33443.jpg")
}