/*
To completely replace the current Go 
process with another. Use exec
*/ 

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Go req's absolute path to binary for exec.
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec needs args in slice form.
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec needs a set of env vars to use.
	env := os.Environ()

	// The syscall
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}