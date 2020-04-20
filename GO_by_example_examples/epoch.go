/*
It is common to get the the time since the Unix epoch.
*/ 

package main

import (
	"fmt"
	"time"
)

func main() {

	// Use time.Now with Unix or UnixNano to get elapsed
	// time since the Unix epoc in sec's or nanosec
	now := time.Now()
	secs := now.Unix()
	nano := now.UnixNano()
	fmt.Println(now)

	// The is no unixMillis, to do this
	// You need to divide from nanosec
	millis := nano / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nano)

	// Convert integer sec's or nanosec's since the epoch
	// into corresponding time
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nano))
}