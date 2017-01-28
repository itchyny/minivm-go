package minivm

import (
	"fmt"
	"os"
	"strconv"
)

func (env *Env) Execute() {
	for env.pc < len(env.code) {
		code := env.code[env.pc]
		switch code.OpCode {
		case OpPrint:
			value := env.stack.Pop()
			fmt.Printf("%v\n", value.Value())
		case OpPop:
			env.stack.Pop()
		case OpDup:
			env.stack.Dup()
		case OpRet:
			v, _ := env.vars.vars[len(env.vars.vars)-1].value.(VInt)
			env.pc = int(v.value)
			env.vars.vars = env.vars.vars[:len(env.vars.vars)-code.Operand]
			env.diffs = env.diffs[:len(env.diffs)-1]
			env.diff -= env.diffs[len(env.diffs)-1]
		case OpCall:
			env.stack.Push(VInt{value: int64(env.pc)})
			fval, _ := env.vars.vars[code.Operand].value.(VFunc)
			reqCap := len(env.vars.vars) + fval.vars
			if reqCap >= cap(env.vars.vars) {
				newVars := make([]Var, reqCap, len(env.vars.vars)+reqCap)
				copy(newVars, env.vars.vars)
				env.vars.vars = newVars
			} else {
				env.vars.vars = env.vars.vars[:reqCap]
			}
			env.diff += env.diffs[len(env.diffs)-1]
			env.diffs = append(env.diffs, fval.vars)
			env.pc = fval.pc
		case OpJmp:
			env.pc += code.Operand
		case OpJmpIf:
			if env.stack.Pop().tobool() {
				env.pc += code.Operand
			}
		case OpJmpNot:
			if !env.stack.Pop().tobool() {
				env.pc += code.Operand
			}
		case OpLetGVar:
			env.vars.vars[code.Operand].value = env.stack.Pop()
		case OpLetLVar:
			env.vars.vars[env.diff+code.Operand].value = env.stack.Pop()
		case OpAdd:
			env.stack.Push(env.stack.Pop().add(env.stack.Pop()))
		case OpSub:
			env.stack.Push(env.stack.Pop().sub(env.stack.Pop()))
		case OpMul:
			env.stack.Push(env.stack.Pop().mul(env.stack.Pop()))
		case OpDiv:
			env.stack.Push(env.stack.Pop().div(env.stack.Pop()))
		case OpGt:
			env.stack.Push(env.stack.Pop().gt(env.stack.Pop()))
		case OpGe:
			env.stack.Push(env.stack.Pop().ge(env.stack.Pop()))
		case OpEq:
			env.stack.Push(env.stack.Pop().eq(env.stack.Pop()))
		case OpNeq:
			env.stack.Push(env.stack.Pop().neq(env.stack.Pop()))
		case OpLt:
			env.stack.Push(env.stack.Pop().lt(env.stack.Pop()))
		case OpLe:
			env.stack.Push(env.stack.Pop().le(env.stack.Pop()))
		case OpPlus:
			env.stack.Push(env.stack.Pop().plus())
		case OpMinus:
			env.stack.Push(env.stack.Pop().minus())
		case OpNot:
			env.stack.Push(env.stack.Pop().not())
		case OpLoadGVar:
			value := env.vars.vars[code.Operand].value
			if value == nil {
				fmt.Fprintln(os.Stderr, "variable not initialized: "+env.vars.vars[code.Operand].name)
				os.Exit(1)
			}
			env.stack.Push(value)
		case OpLoadLVar:
			value := env.vars.vars[env.diff+code.Operand].value
			if value == nil {
				fmt.Fprintln(os.Stderr, "local variable not initialized")
				os.Exit(1)
			}
			env.stack.Push(value)
		case OpLoadT:
			env.stack.Push(VBool{true})
		case OpLoadF:
			env.stack.Push(VBool{false})
		case OpLoad:
			env.stack.Push(env.constant[code.Operand])
		case OpBreak:
			fmt.Fprintln(os.Stderr, "break outside while loop")
			os.Exit(1)
		case OpCont:
			fmt.Fprintln(os.Stderr, "continue outside while loop")
			os.Exit(1)
		default:
			fmt.Fprintln(os.Stderr, "unknown opcode: "+strconv.Itoa(int(code.OpCode)))
			os.Exit(1)
		}
		env.pc++
	}
	if !env.stack.Empty() {
		fmt.Fprintln(os.Stderr, "stack not consumed")
		os.Exit(1)
	}
}
