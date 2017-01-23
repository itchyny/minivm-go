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

%type<node> program expression
%token<token> INT FLOAT

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
	| FLOAT
	{
		value, err := strconv.ParseFloat($1.literal, 64)
		if err != nil {
			yylex.Error("invalid float literal: " + $1.literal)
		}
		$$ = FloatExpr{value: value}
	}

%%
