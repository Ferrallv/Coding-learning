package main

import "fmt"

func tellme() {
	fmt.Println("All my secrets")
}

func main() {
	x := tellme
	x()
}
