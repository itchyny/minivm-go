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
	vars     Vars
}

func Codegen(node Node) *Env {
	env := new(Env)
	env.vars.alloc(node)
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
	case Statements:
		for _, stmt := range node.stmts {
			env.codegen(stmt)
		}
	case LetStmt:
		i := env.vars.lookup(node.ident)
		if i < 0 {
			fmt.Println("unknown variable: " + node.ident)
			os.Exit(1)
		}
		env.codegen(node.expr)
		env.addCode(Code{OpCode: OpLetGVar, Operand: i})
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
	case Ident:
		i := env.vars.lookup(node.name)
		if i < 0 {
			fmt.Println("unknown variable: " + node.name)
			os.Exit(1)
		}
		env.addCode(Code{OpCode: OpLoadGVar, Operand: i})
	case BoolExpr:
		if node.value {
			env.addCode(Code{OpCode: OpLoadT})
		} else {
			env.addCode(Code{OpCode: OpLoadF})
		}
	case IntExpr:
		env.addCode(Code{OpCode: OpLoad, Operand: env.addConst(VInt{node.value})})
	case FloatExpr:
		env.addCode(Code{OpLoad, env.addConst(VFloat{node.value})})
	default:
		fmt.Printf("unknown node type: %+v\n", node)
		os.Exit(1)
	}
}
