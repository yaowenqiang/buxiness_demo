package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommands foo")
		fmt.Println("\tenable:", *fooEnable)
		fmt.Println("\tname:", *fooName)
		fmt.Println("\ttail:", fooCmd.Args())

	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommands bar")
		fmt.Println("\tlevel:", *barLevel)
		fmt.Println("\ttail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
