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
	switch token {
	case scanner.Int:
		token = INT
	case scanner.Float:
		token = FLOAT
	case scanner.Ident:
		switch lexer.scanner.TokenText() {
		case "if":
			token = IF
		case "elseif":
			token = ELSEIF
		case "else":
			token = ELSE
		case "while":
			token = WHILE
		case "break":
			token = BREAK
		case "continue":
			token = CONTINUE
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
	default:
		switch r {
		case '=':
			if lexer.scanner.Peek() == '=' {
				lexer.scanner.Scan()
				token = EQEQ
			} else {
				token = EQ
			}
		case '!':
			if lexer.scanner.Peek() == '=' {
				lexer.scanner.Scan()
				token = NEQ
			} else {
				token = NOT
			}
		case '(':
			token = LPAREN
		case ')':
			token = RPAREN
		case '+':
			token = PLUS
		case '-':
			token = MINUS
		case '*':
			token = TIMES
		case '/':
			token = DIVIDE
		case '>':
			if lexer.scanner.Peek() == '=' {
				lexer.scanner.Scan()
				token = GE
			} else {
				token = GT
			}
		case '<':
			if lexer.scanner.Peek() == '=' {
				lexer.scanner.Scan()
				token = LE
			} else {
				token = LT
			}
		case '&':
			if lexer.scanner.Peek() == '&' {
				lexer.scanner.Scan()
				token = AND
			}
		case '|':
			if lexer.scanner.Peek() == '|' {
				lexer.scanner.Scan()
				token = OR
			}
		case '\r':
			if lexer.scanner.Peek() == '\n' {
				lexer.scanner.Scan()
			}
			token = CR
		case '\n':
			token = CR
		}
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
