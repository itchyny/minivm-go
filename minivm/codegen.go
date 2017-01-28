package minivm

import (
	"fmt"
	"os"
)

func (env *Env) codegen(node Node) {
	switch node := node.(type) {
	case Function:
		i := env.vars.lookup(node.name)
		if i < 0 {
			fmt.Fprintln(os.Stderr, "unknown function name: "+node.name)
			os.Exit(1)
		}
		jmp := env.addCode(Code{OpCode: OpJmp})
		env.localvars = new(Vars)
		env.localvars.allocLocal(node)
		env.localvars.vars = append(env.localvars.vars, Var{})
		env.vars.vars[i].value = VFunc{pc: jmp, vars: len(env.localvars.vars)}
		env.addCode(Code{OpCode: OpLetLVar, Operand: len(env.localvars.vars) - 1})
		for i := len(node.args) - 1; i >= 0; i -= 1 {
			i := env.localvars.lookup(node.args[i])
			env.addCode(Code{OpCode: OpLetLVar, Operand: i})
		}
		env.codegen(node.stmts)
		env.addCode(Code{OpCode: OpLoad, Operand: env.addConst(VInt{0})})
		env.returns = append(env.returns, len(env.code))
		env.addCode(Code{OpCode: OpRet})
		var returns []int
		for _, i := range env.returns {
			if jmp < i && i < len(env.code) {
				env.code[i] = Code{OpCode: OpRet, Operand: len(env.localvars.vars)}
			} else {
				returns = append(returns, i)
			}
		}
		env.returns = returns
		env.localvars = nil
		env.code[jmp].Operand = len(env.code) - jmp - 1
	case ReturnStmt:
		if env.localvars == nil {
			fmt.Fprintln(os.Stderr, "return outside function")
			os.Exit(1)
		}
		if node.expr == nil {
			env.addCode(Code{OpCode: OpLoad, Operand: env.addConst(VInt{0})})
		} else {
			env.codegen(node.expr)
		}
		env.returns = append(env.returns, len(env.code))
		env.addCode(Code{OpCode: OpRet})
	case Statements:
		for _, stmt := range node.stmts {
			env.codegen(stmt)
		}
	case IfStmt:
		env.codegen(node.expr)
		jmpnot := env.addCode(Code{OpCode: OpJmpNot})
		env.codegen(node.stmts)
		if node.elsestmts != nil {
			jmp := env.addCode(Code{OpCode: OpJmp})
			env.code[jmpnot].Operand = len(env.code) - jmpnot - 1
			env.codegen(node.elsestmts)
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
		i := -1
		var local bool
		if env.localvars != nil {
			i = env.localvars.lookup(node.ident)
			local = true
		}
		if i < 0 {
			i = env.vars.lookup(node.ident)
			if i < 0 {
				fmt.Fprintln(os.Stderr, "unknown variable: "+node.ident)
				os.Exit(1)
			}
		}
		env.codegen(node.expr)
		if local {
			env.addCode(Code{OpCode: OpLetLVar, Operand: i})
		} else {
			env.addCode(Code{OpCode: OpLetGVar, Operand: i})
		}
	case PrintStmt:
		env.codegen(node.expr)
		env.addCode(Code{OpCode: OpPrint})
	case CallExpr:
		i := env.vars.lookup(node.name)
		if i < 0 {
			fmt.Fprintln(os.Stderr, "unknown function: "+node.name)
			os.Exit(1)
		}
		for _, expr := range node.exprs {
			env.codegen(expr)
		}
		env.addCode(Code{OpCode: OpCall, Operand: i})
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
	case BinOpExprI:
		env.codegen(node.left)
		var op int8
		switch node.op {
		case PLUS:
			op = OpAddI
		case MINUS:
			op = OpSubI
		case TIMES:
			op = OpMulI
		case DIVIDE:
			op = OpDivI
		case GT:
			op = OpGtI
		case GE:
			op = OpGeI
		case EQEQ:
			op = OpEqI
		case NEQ:
			op = OpNeqI
		case LT:
			op = OpLtI
		case LE:
			op = OpLeI
		default:
			fmt.Fprintln(os.Stderr, "unknown binary operator")
			os.Exit(1)
		}
		env.addCode(Code{OpCode: op, Operand: node.right})
	case UnaryOpExpr:
		env.codegen(node.expr)
		var op int8
		switch node.op {
		case UPLUS:
			op = OpPlus
		case UMINUS:
			op = OpMinus
		case NOT:
			op = OpNot
		default:
			fmt.Fprintln(os.Stderr, "unknown unary operator")
			os.Exit(1)
		}
		env.addCode(Code{OpCode: op})
	case Ident:
		i := -1
		var local bool
		if env.localvars != nil {
			i = env.localvars.lookup(node.name)
			local = true
		}
		if i < 0 {
			i = env.vars.lookup(node.name)
			local = false
			if i < 0 {
				fmt.Fprintln(os.Stderr, "unknown variable: "+node.name)
				os.Exit(1)
			}
		}
		if local {
			env.addCode(Code{OpCode: OpLoadLVar, Operand: i})
		} else {
			env.addCode(Code{OpCode: OpLoadGVar, Operand: i})
		}
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
