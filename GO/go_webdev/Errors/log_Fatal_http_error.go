package main

import (
	"net/http"
	"log"
)

func main() {
	// log.Fatal catches the error that ListenAndServe throws out
	// If it isn't nil, then it prints the error to stdout and gives
	// exit status(1)
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}

// example
func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err :=os.Open("blah.jpg")
	if err != nil {
		// IF we catch an error we right back 
		// the text and the status.
		http.Error(w,. "file not found", 404)
		return 
	}
}