package main

import (
	"fmt"
	"os"
	"io"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := strings.NewReader("Wassup")

	io.Copy(f, r)
}