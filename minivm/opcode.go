package minivm

const (
	OpPrint = iota
	OpPop
	OpDup
	OpRet
	OpCall
	OpJmp
	OpJmpIf
	OpJmpNot
	OpLetGVar
	OpLetLVar
	OpLoadGVar
	OpLoadLVar
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
	OpBreak
	OpCont
)

type Code struct {
	OpCode  int8
	Operand int
}
