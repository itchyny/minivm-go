package minivm

type Node interface{}

type Token struct {
	literal string
	token   int
}

type Statements struct {
	stmts []Node
}

type LetStmt struct {
	ident string
	expr  Node
}

type PrintStmt struct {
	expr Node
}

type BinOpExpr struct {
	op    int
	left  Node
	right Node
}

type Ident struct {
	name string
}

type BoolExpr struct {
	value bool
}

type IntExpr struct {
	value int64
}

type FloatExpr struct {
	value float64
}
