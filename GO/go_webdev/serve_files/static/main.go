package main

import (
	"log"
	"net/http"
)

func main() {
	log.Ftal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}