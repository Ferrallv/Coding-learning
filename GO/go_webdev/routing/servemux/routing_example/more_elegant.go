package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dawg - A DOGGY")
}

type coolcat int

func (c coolcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Me OW me me OW OW - A KITTY")
}

func main() {
	var d hotdog
	var c coolcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat/", c)

	http.ListenAndServe(":8080", mux)
}