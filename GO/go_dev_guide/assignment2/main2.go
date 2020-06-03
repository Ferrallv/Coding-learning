package main

import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
)

func main() {
	// comarg := os.Args[1]

	// file, err := os.Open(comarg)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer file.Close()

	text, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(text))
}