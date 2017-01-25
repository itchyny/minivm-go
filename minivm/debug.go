package minivm

import (
	"fmt"
	"os"
	"strconv"
)

func (env Env) Debug() {
	for i, v := range env.constant {
		fmt.Printf("%d: %v\n", i, v.Value())
	}
	fmt.Println("")

	for i, v := range env.vars.vars {
		fmt.Printf("%d: %s\n", i, v.name)
	}
	fmt.Println("")

	for i, c := range env.code {
		switch c.OpCode {
		case OpPrint:
			fmt.Printf("%d: print\n", i)
		case OpPop:
			fmt.Printf("%d: pop\n", i)
		case OpDup:
			fmt.Printf("%d: dup\n", i)
		case OpJmp:
			fmt.Printf("%d: jmp %d\n", i, c.Operand)
		case OpJmpIf:
			fmt.Printf("%d: jmp_if %d\n", i, c.Operand)
		case OpJmpNot:
			fmt.Printf("%d: jmp_not %d\n", i, c.Operand)
		case OpLetGVar:
			fmt.Printf("%d: let_gvar %d (%s)\n", i, c.Operand, env.vars.vars[c.Operand].name)
		case OpAdd:
			fmt.Printf("%d: add\n", i)
		case OpSub:
			fmt.Printf("%d: sub\n", i)
		case OpMul:
			fmt.Printf("%d: mul\n", i)
		case OpDiv:
			fmt.Printf("%d: div\n", i)
		case OpGt:
			fmt.Printf("%d: gt >\n", i)
		case OpGe:
			fmt.Printf("%d: ge >=\n", i)
		case OpEq:
			fmt.Printf("%d: eq ==\n", i)
		case OpNeq:
			fmt.Printf("%d: neq !=\n", i)
		case OpLt:
			fmt.Printf("%d: lt <\n", i)
		case OpLe:
			fmt.Printf("%d: le <=\n", i)
		case OpNot:
			fmt.Printf("%d: not !\n", i)
		case OpLoadGVar:
			fmt.Printf("%d: load_gvar %d (%s)\n", i, c.Operand, env.vars.vars[c.Operand].name)
		case OpLoadT:
			fmt.Printf("%d: load_true\n", i)
		case OpLoadF:
			fmt.Printf("%d: load_false\n", i)
		case OpLoad:
			fmt.Printf("%d: load %d (%v)\n", i, c.Operand, env.constant[c.Operand].Value())
		case OpBreak:
			fmt.Printf("%d: break\n", i)
		case OpCont:
			fmt.Printf("%d: continue\n", i)
		default:
			fmt.Fprintln(os.Stderr, "unknown opcode: "+strconv.Itoa(int(c.OpCode)))
			os.Exit(1)
		}
	}
}
