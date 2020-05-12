/*
Channels are the pipes that connect concurrent goroutines
You can send values into channels from one goroutine
and receive those values into another goroutine
*/

package main

import "fmt"

func main() {
	// create a new channel with `make(chan val-type`
	// Vhannels are typed by the values they convey
	messages := make(chan string)

	// `Send` a value into a channel using the `channel`
	// syntax. We will send "ping" to the `messages` channel
	// from a new goroutine

	go func() { messages <- "ping"} ()

	// the <- channel syntax receives a value from the channel
	// We'll receive the "ping" and pring it out
	msg := <-messages
	fmt.Println(msg)
}