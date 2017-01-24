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

%type<node> program statements statement expression
%token<token> PRINT
%token<token> EQ LPAREN RPAREN PLUS MINUS TIMES DIVIDE
%token<token> INT FLOAT IDENT CR

%left PLUS MINUS
%left TIMES DIVIDE

%%

program
	: sep_opt statements sep_opt
	{
		$$ = $2
		yylex.(*Lexer).result = $$
	}

statements
	: statement
	{
		$$ = Statements{stmts: []Node{$1}}
	}
	| statements sep statement
	{
		s, _ := $1.(Statements)
		$$ = Statements{stmts: append(s.stmts, $3)}
	}

statement
	: IDENT EQ expression
	{
		$$ = LetStmt{ident: $1.literal, expr: $3}
	}
	| PRINT expression
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
	| LPAREN expression RPAREN
	{
		$$ = $2
	}
	| IDENT
	{
		$$ = Ident{name: $1.literal}
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

sep
	: CR
	| sep CR

sep_opt
	:
	| sep

%%
