package main

import (
	"io"
	"net/http"
)

func landing(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is how we do it!")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "My dog is Bennett")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am Nin")
}

func main() {
	http.HandleFunc("/", landing)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}