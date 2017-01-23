package minivm

const (
	OpPrint = iota
	OpLoad
)

type Code struct {
	OpCode  int8
	Operand int64
}
