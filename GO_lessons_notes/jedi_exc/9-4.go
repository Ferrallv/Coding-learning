package main

import (
	"fmt"
	"sync"
	// "runtime"
)

func main() {

	counter := 0

	gs := 25

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			mu.Unlock() 
			wg.Done()
		}()
	}

	wg.Wait()
}