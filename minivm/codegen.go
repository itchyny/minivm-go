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

func (env *Env) addCode(code Code) int {
	env.code = append(env.code, code)
	return len(env.code) - 1
}

func (env *Env) addConst(value Value) int64 {
	len := len(env.constant)
	env.constant = append(env.constant, value)
	return int64(len)
}

func (env *Env) codegen(node Node) int64 {
	var count int64
	switch node := node.(type) {
	case Statements:
		for _, stmt := range node.stmts {
			count += env.codegen(stmt)
		}
	case IfStmt:
		count += env.codegen(node.expr)
		jmp := env.addCode(Code{OpCode: OpJmpNot})
		count++
		diff := env.codegen(node.stmts)
		env.code[jmp].Operand = diff
		count += diff
	case LetStmt:
		i := env.vars.lookup(node.ident)
		if i < 0 {
			fmt.Fprintln(os.Stderr, "unknown variable: "+node.ident)
			os.Exit(1)
		}
		count += env.codegen(node.expr)
		env.addCode(Code{OpCode: OpLetGVar, Operand: i})
		count++
	case PrintStmt:
		count += env.codegen(node.expr)
		env.addCode(Code{OpCode: OpPrint})
		count++
	case BinOpExpr:
		count += env.codegen(node.left)
		count += env.codegen(node.right)
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
			fmt.Fprintln(os.Stderr, "unknown binary operator")
			os.Exit(1)
		}
		env.addCode(Code{OpCode: op})
		count++
	case Ident:
		i := env.vars.lookup(node.name)
		if i < 0 {
			fmt.Fprintln(os.Stderr, "unknown variable: "+node.name)
			os.Exit(1)
		}
		env.addCode(Code{OpCode: OpLoadGVar, Operand: i})
		count++
	case BoolExpr:
		if node.value {
			env.addCode(Code{OpCode: OpLoadT})
			count++
		} else {
			env.addCode(Code{OpCode: OpLoadF})
			count++
		}
	case IntExpr:
		env.addCode(Code{OpCode: OpLoad, Operand: env.addConst(VInt{node.value})})
		count++
	case FloatExpr:
		env.addCode(Code{OpLoad, env.addConst(VFloat{node.value})})
		count++
	default:
		fmt.Fprintf(os.Stderr, "unknown node type: %+v\n", node)
		os.Exit(1)
	}
	return count
}
