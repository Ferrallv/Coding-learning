package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {

	counter := 0

	gs := 25

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			wg.Done()
		}()
	}

	wg.Wait()
}