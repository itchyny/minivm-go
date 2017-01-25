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
	breaks   []int
	conts    []int
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

func (env *Env) addConst(value Value) int {
	len := len(env.constant)
	env.constant = append(env.constant, value)
	return len
}

func (env *Env) codegen(node Node) {
	switch node := node.(type) {
	case Statements:
		for _, stmt := range node.stmts {
			env.codegen(stmt)
		}
	case IfStmt:
		env.codegen(node.expr)
		jmpnot := env.addCode(Code{OpCode: OpJmpNot})
		env.codegen(node.stmts)
		if node.elsestmts != nil {
			stmts, _ := node.elsestmts.(Statements)
			jmp := env.addCode(Code{OpCode: OpJmp})
			env.code[jmpnot].Operand = len(env.code) - jmpnot - 1
			env.codegen(stmts)
			env.code[jmp].Operand = len(env.code) - jmp - 1
		} else {
			env.code[jmpnot].Operand = len(env.code) - jmpnot - 1
		}
	case WhileStmt:
		pc := len(env.code) - 1
		env.codegen(node.expr)
		jmpnot := env.addCode(Code{OpCode: OpJmpNot})
		env.codegen(node.stmts)
		env.addCode(Code{OpCode: OpJmp, Operand: -(len(env.code) - pc)})
		env.code[jmpnot].Operand = len(env.code) - jmpnot - 1
		var breaks []int
		for _, i := range env.breaks {
			if jmpnot < i && i < len(env.code) {
				env.code[i] = Code{OpCode: OpJmp, Operand: len(env.code) - i - 1}
			} else {
				breaks = append(breaks, i)
			}
		}
		env.breaks = breaks
		var conts []int
		for _, i := range env.conts {
			if jmpnot < i && i < len(env.code) {
				env.code[i] = Code{OpCode: OpJmp, Operand: -(i - pc)}
			} else {
				conts = append(conts, i)
			}
		}
		env.conts = conts
	case BreakStmt:
		env.breaks = append(env.breaks, len(env.code))
		env.addCode(Code{OpCode: OpBreak})
	case ContStmt:
		env.conts = append(env.conts, len(env.code))
		env.addCode(Code{OpCode: OpCont})
	case LetStmt:
		i := env.vars.lookup(node.ident)
		if i < 0 {
			fmt.Fprintln(os.Stderr, "unknown variable: "+node.ident)
			os.Exit(1)
		}
		env.codegen(node.expr)
		env.addCode(Code{OpCode: OpLetGVar, Operand: i})
	case PrintStmt:
		env.codegen(node.expr)
		env.addCode(Code{OpCode: OpPrint})
	case BinOpExpr:
		env.codegen(node.left)
		if node.op == AND {
			env.addCode(Code{OpCode: OpDup})
			jmpnot := env.addCode(Code{OpCode: OpJmpNot})
			env.addCode(Code{OpCode: OpPop})
			env.codegen(node.right)
			env.code[jmpnot].Operand = len(env.code) - jmpnot - 1
		} else if node.op == OR {
			env.addCode(Code{OpCode: OpDup})
			jmpif := env.addCode(Code{OpCode: OpJmpIf})
			env.addCode(Code{OpCode: OpPop})
			env.codegen(node.right)
			env.code[jmpif].Operand = len(env.code) - jmpif - 1
		} else {
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
			case GT:
				op = OpGt
			case GE:
				op = OpGe
			case EQEQ:
				op = OpEq
			case NEQ:
				op = OpNeq
			case LT:
				op = OpLt
			case LE:
				op = OpLe
			default:
				fmt.Fprintln(os.Stderr, "unknown binary operator")
				os.Exit(1)
			}
			env.addCode(Code{OpCode: op})
		}
	case UnaryOpExpr:
		env.codegen(node.expr)
		var op int8
		switch node.op {
		case NOT:
			op = OpNot
		default:
			fmt.Fprintln(os.Stderr, "unknown unary operator")
			os.Exit(1)
		}
		env.addCode(Code{OpCode: op})
	case Ident:
		i := env.vars.lookup(node.name)
		if i < 0 {
			fmt.Fprintln(os.Stderr, "unknown variable: "+node.name)
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
		fmt.Fprintf(os.Stderr, "unknown node type: %+v\n", node)
		os.Exit(1)
	}
}
