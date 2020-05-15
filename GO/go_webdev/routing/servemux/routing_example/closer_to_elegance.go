package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hey, I'm a big woofer")
}

type coolcat int

func (c coolcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hey, I'm a small meower")
}

func main() {
	var d hotdog
	var c coolcat

	http.Handle("/dog/", d)
	http.Handle("/cat/", c)

	http.ListenAndServe(":8080", nil)
}