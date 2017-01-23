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
	env.Execute()
}
