package main

import "fmt"

func main() {
	foo()

	func(){
		fmt.Println("Anonymous func ran")
	}()

	func(x int){
		fmt.Println("What is older than T-Swift feelings.", x)
	}(23)
}

func foo() {
	fmt.Println("foo ran")
}