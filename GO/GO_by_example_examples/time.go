/*
Support for times and durations
*/ 

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Get current time, just like Python!
	now := time.Now()
	p(now)

	// Build a time struct by providing year, month, day, etc.
	// Times are always timezone aware?
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Extract components of hte time value as expected
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Mon-Sun
	p(then.Weekday())

	// Date comparison, between two times. 
	// Tests when the first occurs relative to the second
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Sub method returns a Duration (timedelta?)
	// Interval between two times
	diff := now.Sub(then)
	p(diff)

	// return duration length in different units
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// use Add to advance time, or move backwards
	p(then.Add(diff))
	p(then.Add(-diff))
}