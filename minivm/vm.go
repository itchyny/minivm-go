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
			value := env.stack[len(env.stack)-1]
			fmt.Printf("%v\n", value.Value())
			env.stack = env.stack[:len(env.stack)-1]
		case OpLoad:
			env.stack = append(env.stack, env.constant[code.Operand])
		default:
			fmt.Println("unknown opcode")
			os.Exit(1)
		}
		env.pc++
	}
	if len(env.stack) != 0 {
		fmt.Println("stack not consumed")
		os.Exit(1)
	}
}
