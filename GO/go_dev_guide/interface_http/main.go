package main

import (
	"log"
	"net/http"
	// "os"
	"io"
	"fmt"
)

type LogWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatalln(err)
	}

	lw := LogWriter{}

	io.Copy(lw, resp.Body)
}

func (LogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}