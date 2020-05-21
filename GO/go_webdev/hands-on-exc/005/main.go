// My thought process
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name: "my-cookie",
			Value: "1", 
			Path: "/",
			Visits: 
		})
	} else {
		i, err := strconv.Atoi(c.Value)
		if err != nil {
			fmt.Println(err)
		}
		i += 1
		c.Value = strconv.Itoa(i)

	}

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return 
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}