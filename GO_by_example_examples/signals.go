/*
How to handle signals like SIGTERM and SIGINT gracefully
*/ 

package main 

import (
	"fmt"
	 "os"
	 "os/signal"
	 "syscall"
)

func main() {
	// Go signal notifcation works by sending os.Signal
	// values on a channel.
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// signal.Notify registers the given channel to
	// receive notifications of the specified signal.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// A Goroutine executes a blockign receive for signals
	// It will print the received signal out then notify the 
	// program it can finish.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// The program waits until it gets the expected signal
	// then it exits.
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}