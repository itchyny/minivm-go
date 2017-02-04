package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/itchyny/minivm-go/minivm"
)

var debug = flag.Bool("debug", false, "debug code instructions")

func main() {
	flag.Parse()
	lexer := new(minivm.Lexer)
	lexer.Init(os.Stdin)
	minivm.Parse(lexer)
	env, err := minivm.Codegen(lexer.Result())
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if debug != nil && *debug {
		env.Debug()
	}
	env.Execute()
}
