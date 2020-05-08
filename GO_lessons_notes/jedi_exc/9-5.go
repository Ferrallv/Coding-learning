package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64 = 0

	gs := 25

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
}
