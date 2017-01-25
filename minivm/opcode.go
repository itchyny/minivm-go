package minivm

const (
	OpPrint = iota
	OpPop
	OpDup
	OpJmp
	OpJmpIf
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
	OpGt
	OpGe
	OpEq
	OpNeq
	OpLt
	OpLe
	OpNot

	// temporary opcodes
	OpBreak
	OpCont
)

type Code struct {
	OpCode  int8
	Operand int
}
