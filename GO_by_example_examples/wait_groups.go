/*
To wait for multiple goroutines to finsih, we can use a 
wait group
*/ 

package main

import (
	"fmt"
	"sync"
	"time"
)

// this is the function we'll run in every goroutine. Note
// that a WaitGroup must be passed to functions by pointer
func worker(id int, wg * sync.WaitGroup) {

	// on return, notify the WaitGroup that we're done.
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// Sleep to simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// This WaitGrop is used to wait for all the goroutines
	// launched here to finish
	var wg sync.WaitGroup

	// Launch severl goroutines and increment the WaitGroup 
	// counter for each
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the WaitGroup couanter goes back to 0;
	// all the workers notfied they're done.
	wg.Wait()
}