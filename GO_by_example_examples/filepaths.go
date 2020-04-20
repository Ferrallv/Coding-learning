/*
The filepath package provides functions to parse
and construct file paths in a way that is protable between
operating systems; dir/file on linux vs dir\file on 
Windows.
*/ 

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// Join should be used ot construct paths in a portable way
	// It takes any number of arguments and constructs a hierarchical
	// path from them
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Always use Join instead of concatenatin /s or \s manually.
	// Join will normalize paths by removing superfluous seps and dir changes
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1/", "filename"))

	// Dir and Base can be used to split a path to the directory
	// and the file. Alternatively, Split will return both 
	// in the same call.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// To check if a path is absolute
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Some file names have extensions following a dot
	// Split the extension with Ext
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Finding a files name with extension removed
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel finds a relative path between a base and a target.
	// If the target cannot be made relative to base, it returns
	// an error
	rel, err := filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}