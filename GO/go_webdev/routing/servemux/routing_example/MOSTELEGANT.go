package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Holy crap I'm a beautiful dog")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I KNOW I am the prettiest cat")
}

func main() {
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat/", c)

	http.ListenAndServe(":8080", nil)
}