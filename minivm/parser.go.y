%{
package minivm

import "strconv"

func Parse(yylex yyLexer) int {
	return yyParse(yylex)
}
%}

%union{
	node  Node
	statements Statements
	literals []string
	nodes []Node
	token Token
}

%type<node> program statement else_opt expression expression_opt
%type<statements> statements
%type<literals> fargs farg_list
%type<nodes> args arg_list
%token<token> FUNC RETURN IF ELSEIF ELSE WHILE BREAK CONTINUE END PRINT
%token<token> EQ LPAREN RPAREN COMMA
%token<token> PLUS MINUS TIMES DIVIDE UPLUS UMINUS
%token<token> GT GE EQEQ NEQ LT LE NOT
%token<token> INT FLOAT TRUE FALSE IDENT CR

%left OR
%left AND
%nonassoc EQEQ NEQ
%left GT GE LT LE
%left PLUS MINUS
%left TIMES DIVIDE
%right NOT UPLUS UMINUS

%%

program
	: statements
	{
		$$ = $1
		yylex.(*Lexer).result = $$
	}

statements
	:
	{
		$$ = Statements{stmts: []Node{}}
	}
	| statements statement sep
	{
		$$ = Statements{stmts: append($1.stmts, $2)}
	}

statement
	: FUNC IDENT LPAREN fargs RPAREN sep statements END
	{
		$$ = Function{name: $2.literal, args: $4, stmts: $7}
	}
	| RETURN expression_opt
	{
		$$ = ReturnStmt{expr: $2}
	}
	| IF expression sep statements else_opt END
	{
		$$ = IfStmt{expr: $2, stmts: $4, elsestmts: $5}
	}
	| WHILE expression sep statements END
	{
		$$ = WhileStmt{expr: $2, stmts: $4}
	}
	| BREAK
	{
		$$ = BreakStmt{}
	}
	| CONTINUE
	{
		$$ = ContStmt{}
	}
	| IDENT EQ expression
	{
		$$ = LetStmt{ident: $1.literal, expr: $3}
	}
	| PRINT expression
	{
		$$ = PrintStmt{expr: $2}
	}

expression_opt
	:
	{
		$$ = nil
	}
	| expression
	{
		$$ = $1
	}

fargs
	:
	{
		$$ = []string{}
	}
	| farg_list
	{
		$$ = $1
	}

farg_list
	: IDENT
	{
		$$ = []string{$1.literal}
	}
	| farg_list COMMA IDENT
	{
		$$ = append($1, $3.literal)
	}

args
	:
	{
		$$ = []Node{}
	}
	| arg_list
	{
		$$ = $1
	}

arg_list
	: expression
	{
		$$ = []Node{$1}
	}
	| arg_list COMMA expression
	{
		$$ = append($1, $3)
	}

else_opt
	:
	{
		$$ = nil
	}
	| ELSEIF expression sep statements else_opt
	{
		$$ = Statements{stmts: []Node{IfStmt{expr: $2, stmts: $4, elsestmts: $5}}}
	}
	| ELSE sep statements
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
	| expression OR expression
	{
		$$ = BinOpExpr{op: OR, left: $1, right: $3}
	}
	| expression AND expression
	{
		$$ = BinOpExpr{op: AND, left: $1, right: $3}
	}
	| PLUS expression %prec UPLUS
	{
		$$ = UnaryOpExpr{op: UPLUS, expr: $2}
	}
	| MINUS expression %prec UMINUS
	{
		$$ = UnaryOpExpr{op: UMINUS, expr: $2}
	}
	| NOT expression
	{
		$$ = UnaryOpExpr{op: NOT, expr: $2}
	}
	| LPAREN expression RPAREN
	{
		$$ = $2
	}
	| IDENT LPAREN args RPAREN
	{
		$$ = CallExpr{name: $1.literal, exprs: $3}
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

%%
