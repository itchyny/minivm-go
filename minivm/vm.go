package minivm

import (
	"fmt"
	"os"
	"strconv"
)

func (env Env) Execute() {
	for env.pc < len(env.code) {
		code := env.code[env.pc]
		switch code.OpCode {
		case OpPrint:
			value := env.stack.Pop()
			fmt.Printf("%v\n", value.Value())
		case OpJmp:
			env.pc += code.Operand
		case OpJmpNot:
			if !env.stack.Pop().tobool() {
				env.pc += code.Operand
			}
		case OpLetGVar:
			env.vars.vars[code.Operand].value = env.stack.Pop()
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
		case OpLoadGVar:
			value := env.vars.vars[code.Operand].value
			if value == nil {
				fmt.Fprintln(os.Stderr, "variable not initialized: "+env.vars.vars[code.Operand].name)
				os.Exit(1)
			}
			env.stack.Push(value)
		case OpLoadT:
			env.stack.Push(VBool{true})
		case OpLoadF:
			env.stack.Push(VBool{false})
		case OpLoad:
			env.stack.Push(env.constant[code.Operand])
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
