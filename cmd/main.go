package main

import (
	"flag"
	"os"

	"github.com/itchyny/minivm-go/minivm"
)

var debug = flag.Bool("debug", false, "debug code instructions")

func main() {
	flag.Parse()
	lexer := new(minivm.Lexer)
	lexer.Init(os.Stdin)
	minivm.Parse(lexer)
	env := minivm.Codegen(lexer.Result())
	if debug != nil && *debug {
		env.Debug()
	}
	env.Execute()
}
