/*
Reading and writing files
*/ 

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Slurping a file's entire contents into memory
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// With more control on how and what to read
	f, err := os.Open("/tmp/dat")
	check(err)

	// Read the first few bytes.
	// Allow up to 5, but note how many we read.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// You can also Seek to a known location
	// and Read from there!
	o2, err := f.Seek(6,0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes: @ %d\n", n2, o2)	
	fmt.Printf("%v\n", string(b2[:n2]))

	// The io package provides som helpful file reading funcs
	// repeat above with ReadAtLeast
	o3, err := f.Seek(6,0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes: @ %d: %s\n", n3, o3, string(b3))	
	
	// No built-in for rewind, but Seek(0, 0) does this
	_, err = f.Seek(0, 0)
	check(err)

	// bufio package implements a buffered reader that may
	// be useful for it's efficiency with small reads
	// and additional reading methods it provides
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Close the file when you are done! You can alternatively
	// schedule this immediately after Opening with defer.
	f.Close()
}