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
	OpAddI
	OpSubI
	OpMulI
	OpDivI
	OpGtI
	OpGeI
	OpEqI
	OpNeqI
	OpLtI
	OpLeI
	OpPlus
	OpMinus
	OpNot
	OpBreak
	OpCont
)

type Code struct {
	OpCode  int8
	Operand int
}
