package minivm

const (
	OpPrint = iota
	OpLoad
	OpAdd
	OpSub
	OpMul
	OpDiv
)

type Code struct {
	OpCode  int8
	Operand int64
}
