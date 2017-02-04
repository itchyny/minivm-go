package minivm

import (
	"os"
)

func Execute(filename string, debug bool) error {
	lexer := new(Lexer)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	lexer.Init(file)
	Parse(lexer)
	env, err := Codegen(lexer.Result())
	if err != nil {
		return err
	}
	if debug {
		env.Debug()
	}
	env.Execute()
	return nil
}
