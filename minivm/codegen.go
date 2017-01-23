package minivm

import (
	"fmt"
	"os"
)

type Env struct {
	pc       int
	code     []Code
	constant []Value
	stack    []Value
}

func Codegen(node Node) *Env {
	env := new(Env)
	env.codegen(node)
	return env
}

func (env *Env) addCode(code Code) int64 {
	env.code = append(env.code, code)
	return 1
}

func (env *Env) addConst(value Value) int64 {
	len := len(env.constant)
	env.constant = append(env.constant, value)
	return int64(len)
}

func (env *Env) codegen(node Node) {
	switch node := node.(type) {
	case IntExpr:
		env.addCode(Code{OpCode: OpLoad, Operand: env.addConst(VInt{node.value})})
		env.addCode(Code{OpCode: OpPrint})
	case FloatExpr:
		env.addCode(Code{OpLoad, env.addConst(VFloat{node.value})})
		env.addCode(Code{OpCode: OpPrint})
	default:
		fmt.Println("unknown node type")
		os.Exit(1)
	}
}
