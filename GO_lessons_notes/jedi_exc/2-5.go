package main

import "fmt"

func main() {
	s := `This is a RAW STRING! 
	Not safe to eat`
	sl := fmt.Sprint(s)

	fmt.Println(sl)

}