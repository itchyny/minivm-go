%{
package minivm

import "strconv"

func Parse(yylex yyLexer) int {
	return yyParse(yylex)
}
%}

%union{
	node  Node
	token Token
}

%type<node> program statement expression
%token<token> PRINT
%token<token> PLUS MINUS TIMES DIVIDE
%token<token> INT FLOAT

%left PLUS MINUS
%left TIMES DIVIDE

%%

program
	: statement
	{
		$$ = $1
		yylex.(*Lexer).result = $$
	}

statement
	: PRINT expression
	{
		$$ = PrintStmt{expr: $2}
	}

expression
	: expression PLUS expression
	{
		$$ = BinOpExpr{op: PLUS, left: $1, right: $3}
	}
	| expression MINUS expression
	{
		$$ = BinOpExpr{op: MINUS, left: $1, right: $3}
	}
	| expression TIMES expression
	{
		$$ = BinOpExpr{op: TIMES, left: $1, right: $3}
	}
	| expression DIVIDE expression
	{
		$$ = BinOpExpr{op: DIVIDE, left: $1, right: $3}
	}
	| INT
	{
		value, err := strconv.ParseInt($1.literal, 10, 64)
		if err != nil {
			yylex.Error("invalid integer literal: " + $1.literal)
		}
		$$ = IntExpr{value: value}
	}
	| FLOAT
	{
		value, err := strconv.ParseFloat($1.literal, 64)
		if err != nil {
			yylex.Error("invalid float literal: " + $1.literal)
		}
		$$ = FloatExpr{value: value}
	}

%%
