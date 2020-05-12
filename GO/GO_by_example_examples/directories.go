 /*
Useful Go functions for working with dir's
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

func main () {
	// Create a new sub-dir
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Good practice to defer the removal of temp dir's
	// similar to rm -rf
	defer os.RemoveAll("subdir")

	// Helper func to create new empty file
	createEmptyFile := func(name string) {
	 	 d := []byte("")
	 	 check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// Create a hierarchy of dir's, including parents.
	// similar to mkdir -p
	err = os.MkdirAll("subdir/parent/child", 0755)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir lists dir contents, returnong a slice
	// of os.FileInfo objects
	c, err := ioutil.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Chdir lets us change the current working dir, like cd
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Now we'll see the contents of subdir/parent/child
	// when we list the current dir
	c, err = ioutil.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// cd to where we started
	err = os.Chdir("../../..")
	check(err)

	// Visit a dir recursively. Walk accepts a callback func
	// to handle every file or dir visited
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// Visit is called for every file or dir found recursively
// by filepath.Walk
func visit (p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}