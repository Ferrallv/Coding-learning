package main

import (
	"fmt"
	"sync"
	"runtime"
)


// My answer!
// func main() {

// 	counter := 0
// 	const bells = 50

// 	var wg sync.WaitGroup
// 	wg.Add(bells)

// 	var mu sync.Mutex

// 	for i := 0; i <= bells; i++ {
// 		go func() {
// 			mu.Lock()
// 			fmt.Println(counter)
// 			counter++
// 			mu.Unlock()
// 			wg.Done()
// 		}()
// 		fmt.Println("Routines:  ", runtime.NumGoroutine())
// 	}

// 	wg.Wait()
// }

// Todds answer
func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Begin CPU", runtime.NumCPU())
	fmt.Println("Begin gs", runtime.NumGoroutine())

	go func () {
		fmt.Println("I am first anon Function, hello!")
		wg.Done()
	}()

	go func () {
		fmt.Println("I am second anon Function, hello!")
		wg.Done()
	}()

	fmt.Println("Mid CPU", runtime.NumCPU())
	fmt.Println("Mid gs", runtime.NumGoroutine())
	
	wg.Wait()

	fmt.Println("End CPU", runtime.NumCPU())
	fmt.Println("End gs", runtime.NumGoroutine())

	fmt.Println("Gonna exit now, thanks!")
}