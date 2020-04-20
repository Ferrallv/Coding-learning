/*
We can make our own subcommands
*/ 

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Dec subcommands using NewFlagSet.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Define different supported flags
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Subcommand is expected as the first arg to the program
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Check which subcommand is invoked
	switch os.Args[1] {

		// For every subcommand, parse its own flag
		// and have access to trailing positional args
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("   enable:", *fooEnable)
		fmt.Println("   name:", *fooName)
		fmt.Println("   tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("   level:", *barLevel)
		fmt.Println("   tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}