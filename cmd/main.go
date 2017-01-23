package main

import (
	"fmt"
	"os"

	"github.com/itchyny/minivm-go/minivm"
)

func main() {
	lexer := new(minivm.Lexer)
	lexer.Init(os.Stdin)
	minivm.Parse(lexer)
	fmt.Printf("%+v\n", lexer.Result())
}
