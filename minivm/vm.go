package minivm

import (
	"fmt"
	"os"
)

func (env Env) Execute() {
	for env.pc < len(env.code) {
		code := env.code[env.pc]
		switch code.OpCode {
		case OpPrint:
			value := env.stack.Pop()
			fmt.Printf("%v\n", value.Value())
		case OpLoad:
			env.stack.Push(env.constant[code.Operand])
		default:
			fmt.Println("unknown opcode")
			os.Exit(1)
		}
		env.pc++
	}
	if !env.stack.Empty() {
		fmt.Println("stack not consumed")
		os.Exit(1)
	}
}
