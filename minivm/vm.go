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
		case OpJmpNot:
			if !env.stack.Pop().tobool() {
				env.pc += int(code.Operand)
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
		case OpLoadGVar:
			env.stack.Push(env.vars.vars[code.Operand].value)
		case OpLoadT:
			env.stack.Push(VBool{true})
		case OpLoadF:
			env.stack.Push(VBool{false})
		case OpLoad:
			env.stack.Push(env.constant[code.Operand])
		default:
			fmt.Println("unknown opcode: " + strconv.Itoa(int(code.OpCode)))
			os.Exit(1)
		}
		env.pc++
	}
	if !env.stack.Empty() {
		fmt.Println("stack not consumed")
		os.Exit(1)
	}
}
