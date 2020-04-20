
/*
Spawning processes in Go
*/ 

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// exec.Command helper creates an object 
	// to represent a command that prints something to stdout
	dateCmd := exec.Command("date")

	// .Output will help us handle the common case fo running
	// a command, waiting for it to finish, and collecting
	// it's output
	// dateOut will hold bytes with date info, if no erorrs
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Pipe data to the external process on its stdin
	// Collect results for its stdout.
	grepCmd := exec.Command("grep", "hello")

	// Grab input/output pipes, start the process, 
	// write some input to it, read the resulting output
	// wait for process to end.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))
	// When spawning commands, you must provide an explicitly
	// delineated command and argument array, vs. being able to just
	// pass in one command-line string. To spawn a full command 
	// with string, you can use bash's -c option.
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}