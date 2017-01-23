%{
package minivm

import (
	"strconv"
)

type Node interface{}

type Token struct {
	literal string
	token   int
}

type IntExpr struct {
	value int64
}

func Parse(yylex yyLexer) int {
	return yyParse(yylex)
}
%}

%union{
	node  Node
	token Token
}

%type<node> program expression
%token<token> INT

%%

program
	: expression
	{
		$$ = $1
		yylex.(*Lexer).result = $$
	}

expression
	: INT
	{
		value, err := strconv.ParseInt($1.literal, 10, 64)
		if err != nil {
			yylex.Error("invalid integer literal: " + $1.literal)
		}
		$$ = IntExpr{value: value}
	}

%%
