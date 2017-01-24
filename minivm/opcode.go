package minivm

const (
	OpPrint = iota
	OpJmpNot
	OpLetGVar
	OpLoadGVar
	OpLoadT
	OpLoadF
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
