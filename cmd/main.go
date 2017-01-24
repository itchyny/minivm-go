package main

import (
	"os"

	"github.com/itchyny/minivm-go/minivm"
)

func main() {
	lexer := new(minivm.Lexer)
	lexer.Init(os.Stdin)
	minivm.Parse(lexer)
	env := minivm.Codegen(lexer.Result())
	if len(os.Args) > 1 && os.Args[1] == "--debug" {
		env.Debug()
	}
	env.Execute()
}
