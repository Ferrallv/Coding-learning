/*
Defer is used to ensure that a function call is performed
later in a program's execution, usually for cleanup.
defer is often used where ensure(?) and finally(?) 
would be used in other languages.
*/ 

package main
import (
	"fmt"
	"os"
)

// To create a file, write to it, and then close it.
func main() {
	// Immediately after getting a file object with createFile,
	// we defer the closing of that file with closeFile. This
	// will be executed at the end of the enclosing function (main)
	// after writeFile has finished.

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile (f)
}

func createFile(p string) *os.File {
	fmt.Println("Creating our nice little file")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "YO, I GOTTA LOTTA DATA IN ME!")
}

// It's important to check for errors when closing
func closeFile(f *os.File) {
	fmt.Println("Closing the created file, defferedly")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}