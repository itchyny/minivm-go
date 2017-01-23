package minivm

import (
	"fmt"
	"os"
)

type Env struct {
	pc       int
	code     []Code
	constant []Value
	stack    Stack
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
	case PrintStmt:
		env.codegen(node.expr)
		env.addCode(Code{OpCode: OpPrint})
	case BinOpExpr:
		env.codegen(node.left)
		env.codegen(node.right)
		var op int8
		switch node.op {
		case PLUS:
			op = OpAdd
		case MINUS:
			op = OpSub
		case TIMES:
			op = OpMul
		case DIVIDE:
			op = OpDiv
		default:
			fmt.Println("unknown binary operator")
			os.Exit(1)
		}
		env.addCode(Code{OpCode: op})
	case IntExpr:
		env.addCode(Code{OpCode: OpLoad, Operand: env.addConst(VInt{node.value})})
	case FloatExpr:
		env.addCode(Code{OpLoad, env.addConst(VFloat{node.value})})
	default:
		fmt.Println("unknown node type")
		os.Exit(1)
	}
}
