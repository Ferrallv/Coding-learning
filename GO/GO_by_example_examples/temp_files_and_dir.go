/*
Temp files and dir's
*/ 

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// ioutil.TempFile is the easiest to create temp files
	f, err := ioutil.TempFile("", "sample")
	check(err)

	// Display name of temp file. Done automatically
	// using the prefix assigned in ioutil.TempFile
	fmt.Println("Temp file name:", f.Name())

	// Clean up the file when we're done
	defer os.Remove(f.Name())

	// Write some data
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// If we're writing a bunch of temp files a temp dir
	// might be preferable. Use ioutil.TempDir
	dname, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// Synth the temp file names by prefixing them with
	// temp dir.
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}