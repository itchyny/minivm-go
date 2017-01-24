package minivm

const (
	OpPrint = iota
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
