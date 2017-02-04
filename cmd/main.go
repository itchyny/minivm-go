package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/itchyny/minivm-go/minivm"
)

var debugFlag = flag.Bool("debug", false, "debug code instructions")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Specify one filename\n")
		os.Exit(1)
	}
	debug := false
	if debugFlag != nil {
		debug = *debugFlag
	}
	if err := minivm.Execute(args[0], debug); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
