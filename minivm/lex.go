package minivm

import (
	"fmt"
	"io"
	"os"
	"text/scanner"
)

type Lexer struct {
	scanner scanner.Scanner
	result  Node
}

func (lexer *Lexer) Init(reader io.Reader) {
	lexer.scanner.Init(reader)
}

func (lexer *Lexer) Lex(lval *yySymType) int {
	r := lexer.scanner.Scan()
	token := int(r)
	if token == scanner.Int {
		token = INT
	} else if token == scanner.Float {
		token = FLOAT
	} else if token == scanner.Ident {
		switch lexer.scanner.TokenText() {
		case "print":
			token = PRINT
		}
	} else if r == '+' {
		token = PLUS
	} else if r == '-' {
		token = MINUS
	} else if r == '*' {
		token = TIMES
	} else if r == '/' {
		token = DIVIDE
	}
	lval.token = Token{token: token, literal: lexer.scanner.TokenText()}
	return token
}

func (lexer *Lexer) Error(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func (lexer *Lexer) Result() Node {
	return lexer.result
}
