package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	subcmdSeed := newSubcmdSeed()

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Printf(
			"one of subcommands: %v is required\n",
			[]string{
				subcmdSeed.getCmdName(),
			},
		)
		os.Exit(1)
	}

	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	switch os.Args[1] {
	case subcmdSeed.getCmdName():
		if err := subcmdSeed.parse(os.Args[2:]); err != nil {
			fmt.Printf("error in cmd %s: %+v", subcmdSeed.getCmdName(), err)
			subcmdSeed.printDefaults()
			os.Exit(1)
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
