package minivm

type Node interface{}

type Token struct {
	literal string
	token   int
}

type IntExpr struct {
	value int64
}

type FloatExpr struct {
	value float64
}
