package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go CheckLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			CheckLink(link, c)
		}(l)
	}
}

func CheckLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		// fmt.Println(link, "might be down!")
		c <- link + "Might be down."
		return
	}
	
	fmt.Println(link, "is up!")
	c <- link + " is up"
}
