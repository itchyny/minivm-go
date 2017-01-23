package minivm

import (
	"fmt"
	"os"
	"text/scanner"
)

type Lexer struct {
	scanner.Scanner
	result Node
}

func (lexer *Lexer) Lex(lval *yySymType) int {
	token := int(lexer.Scan())
	if token == scanner.Int {
		token = INT
	} else if token == scanner.Float {
		token = FLOAT
	}
	lval.token = Token{token: token, literal: lexer.TokenText()}
	return token
}

func (lexer *Lexer) Error(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func (lexer *Lexer) Result() Node {
	return lexer.result
}
