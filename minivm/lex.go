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
	lexer.scanner.Whitespace = 1<<'\t' | 1<<' '
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
		case "if":
			token = IF
		case "elseif":
			token = ELSEIF
		case "else":
			token = ELSE
		case "end":
			token = END
		case "print":
			token = PRINT
		case "true":
			token = TRUE
		case "false":
			token = FALSE
		default:
			token = IDENT
		}
	} else if r == '=' {
		token = EQ
	} else if r == '(' {
		token = LPAREN
	} else if r == ')' {
		token = RPAREN
	} else if r == '+' {
		token = PLUS
	} else if r == '-' {
		token = MINUS
	} else if r == '*' {
		token = TIMES
	} else if r == '/' {
		token = DIVIDE
	} else if r == '\r' {
		if lexer.scanner.Peek() == '\n' {
			lexer.scanner.Scan()
		}
		token = CR
	} else if r == '\n' {
		token = CR
	}
	lval.token = Token{token: token, literal: lexer.scanner.TokenText()}
	return token
}

func (lexer *Lexer) Error(err string) {
	fmt.Fprintln(os.Stderr, lexer.scanner.Pos().String()+": "+err)
	os.Exit(1)
}

func (lexer *Lexer) Result() Node {
	return lexer.result
}
