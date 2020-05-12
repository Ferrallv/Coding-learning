// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // My answer ends with a fatal error. I don't know
// // where to close c.
// func main() {
// 	c := make(chan int)

// 	var gs = 10

// 	var wg sync.WaitGroup
// 	wg.Add(gs)


// 	for i := 0; i < gs ; i++ {
// 		go func() {
// 			for i := 0; i < 10; i++ {
// 				c <- i
// 			}
// 			wg.Done()
// 		}()
// 	}


	
// 	for v := range c {
// 		fmt.Println(v)
// 	}

// 	wg.Wait()
// }

// Todds answer

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")
}