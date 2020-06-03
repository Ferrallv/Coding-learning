package main

import (
	"log"
	"net/http"
	"os"
	"io"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatalln(err)
	}

	io.Copy(os.Stdout, resp.Body)
}