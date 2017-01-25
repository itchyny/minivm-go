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

%type<node> program statements statement else_opt expression
%token<token> IF ELSEIF ELSE END PRINT
%token<token> EQ LPAREN RPAREN
%token<token> PLUS MINUS TIMES DIVIDE
%token<token> GT GE EQEQ NEQ LT LE
%token<token> INT FLOAT TRUE FALSE IDENT CR

%nonassoc EQEQ NEQ
%left GT GE LT LE
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
	: IF expression sep statements sep else_opt END
	{
		$$ = IfStmt{expr: $2, stmts: $4, elsestmts: $6}
	}
	| IDENT EQ expression
	{
		$$ = LetStmt{ident: $1.literal, expr: $3}
	}
	| PRINT expression
	{
		$$ = PrintStmt{expr: $2}
	}

else_opt
	:
	{
		$$ = nil
	}
	| ELSEIF expression sep statements sep else_opt
	{
		$$ = Statements{stmts: []Node{IfStmt{expr: $2, stmts: $4, elsestmts: $6}}}
	}
	| ELSE sep statements sep
	{
		$$ = $3
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
	| expression GT expression
	{
		$$ = BinOpExpr{op: GT, left: $1, right: $3}
	}
	| expression GE expression
	{
		$$ = BinOpExpr{op: GE, left: $1, right: $3}
	}
	| expression EQEQ expression
	{
		$$ = BinOpExpr{op: EQEQ, left: $1, right: $3}
	}
	| expression NEQ expression
	{
		$$ = BinOpExpr{op: NEQ, left: $1, right: $3}
	}
	| expression LT expression
	{
		$$ = BinOpExpr{op: LT, left: $1, right: $3}
	}
	| expression LE expression
	{
		$$ = BinOpExpr{op: LE, left: $1, right: $3}
	}
	| LPAREN expression RPAREN
	{
		$$ = $2
	}
	| IDENT
	{
		$$ = Ident{name: $1.literal}
	}
	| TRUE
	{
		$$ = BoolExpr{value: true}
	}
	| FALSE
	{
		$$ = BoolExpr{value: false}
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
