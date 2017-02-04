package minivm

import (
	"errors"
	"fmt"
)

func (env *Env) codegen(node Node) (int, error) {
	switch node := node.(type) {
	case Function:
		if env.localvars != nil {
			return VTUnknown, errors.New("you cannot define a function in a function: " + node.name)
		}
		i := env.vars.lookup(node.name)
		if i < 0 {
			return VTUnknown, errors.New("unknown function name: " + node.name)
		}
		jmp := env.addCode(Code{OpCode: OpJmp})
		env.localvars = new(Vars)
		if err := env.localvars.allocLocal(node); err != nil {
			return VTUnknown, err
		}
		env.localvars.vars = append(env.localvars.vars, Var{})
		env.vars.vars[i].value = VFunc{pc: jmp, vars: len(env.localvars.vars)}
		env.addCode(Code{OpCode: OpLetLVar, Operand: len(env.localvars.vars) - 1})
		for i := len(node.args) - 1; i >= 0; i -= 1 {
			j := env.localvars.lookup(node.args[i])
			env.addCode(Code{OpCode: OpLetLVar, Operand: j})
		}
		if vtype, err := env.codegen(node.stmts); err != nil {
			return vtype, err
		}
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
			return VTUnknown, errors.New("return outside function")
		}
		if node.expr == nil {
			env.addCode(Code{OpCode: OpLoad, Operand: env.addConst(VInt{0})})
		} else {
			if vtype, err := env.codegen(node.expr); err != nil {
				return vtype, err
			}
		}
		env.returns = append(env.returns, len(env.code))
		env.addCode(Code{OpCode: OpRet})
	case Statements:
		for _, stmt := range node.stmts {
			if vtype, err := env.codegen(stmt); err != nil {
				return vtype, err
			}
		}
	case IfStmt:
		if vtype, err := env.codegen(node.expr); err != nil {
			return vtype, err
		}
		jmpnot := env.addCode(Code{OpCode: OpJmpNot})
		if vtype, err := env.codegen(node.stmts); err != nil {
			return vtype, err
		}
		if node.elsestmts != nil {
			jmp := env.addCode(Code{OpCode: OpJmp})
			env.code[jmpnot].Operand = len(env.code) - jmpnot - 1
			if vtype, err := env.codegen(node.elsestmts); err != nil {
				return vtype, err
			}
			env.code[jmp].Operand = len(env.code) - jmp - 1
		} else {
			env.code[jmpnot].Operand = len(env.code) - jmpnot - 1
		}
	case WhileStmt:
		pc := len(env.code) - 1
		if vtype, err := env.codegen(node.expr); err != nil {
			return vtype, err
		}
		jmpnot := env.addCode(Code{OpCode: OpJmpNot})
		if vtype, err := env.codegen(node.stmts); err != nil {
			return vtype, err
		}
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
				return VTUnknown, errors.New("unknown variable: " + node.ident)
			}
			local = false
		}
		vtype, err := env.codegen(node.expr)
		if err != nil {
			return vtype, err
		}
		if local {
			env.addCode(Code{OpCode: OpLetLVar, Operand: i})
			env.localvars.vars[i].vtype = vtype
		} else {
			env.addCode(Code{OpCode: OpLetGVar, Operand: i})
			env.vars.vars[i].vtype = vtype
		}
	case PrintStmt:
		if vtype, err := env.codegen(node.expr); err != nil {
			return vtype, err
		}
		env.addCode(Code{OpCode: OpPrint})
	case CallExpr:
		i := -1
		var local bool
		if env.localvars != nil {
			i = env.localvars.lookup(node.name)
			local = true
		}
		if i < 0 {
			i = env.vars.lookup(node.name)
			if i < 0 {
				return VTUnknown, errors.New("unknown function: " + node.name)
			}
			local = false
		}
		for _, expr := range node.exprs {
			if vtype, err := env.codegen(expr); err != nil {
				return vtype, err
			}
		}
		if local {
			env.addCode(Code{OpCode: OpCallL, Operand: i})
		} else {
			env.addCode(Code{OpCode: OpCallG, Operand: i})
		}
	case BinOpExpr:
		if vtype, err := env.codegen(node.left); err != nil {
			return vtype, err
		}
		if node.op == AND {
			env.addCode(Code{OpCode: OpDup})
			jmpnot := env.addCode(Code{OpCode: OpJmpNot})
			env.addCode(Code{OpCode: OpPop})
			if vtype, err := env.codegen(node.right); err != nil {
				return vtype, err
			}
			env.code[jmpnot].Operand = len(env.code) - jmpnot - 1
		} else if node.op == OR {
			env.addCode(Code{OpCode: OpDup})
			jmpif := env.addCode(Code{OpCode: OpJmpIf})
			env.addCode(Code{OpCode: OpPop})
			if vtype, err := env.codegen(node.right); err != nil {
				return vtype, err
			}
			env.code[jmpif].Operand = len(env.code) - jmpif - 1
		} else {
			if vtype, err := env.codegen(node.right); err != nil {
				return vtype, err
			}
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
				return VTUnknown, errors.New("unknown binary operator")
			}
			env.addCode(Code{OpCode: op})
		}
	case BinOpExprI:
		if vtype, err := env.codegen(node.left); err != nil {
			return vtype, err
		}
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
			return VTUnknown, errors.New("unknown binary operator")
		}
		env.addCode(Code{OpCode: op, Operand: node.right})
	case UnaryOpExpr:
		vtype, err := env.codegen(node.expr)
		if err != nil {
			return vtype, err
		}
		var op int8
		switch node.op {
		case UPLUS:
			if vtype != VTUnknown && vtype != VTInt && vtype != VTFloat {
				return vtype, errors.New("invalid unary operator + on type: " + VTString(vtype))
			}
			op = OpPlus
		case UMINUS:
			if vtype != VTUnknown && vtype != VTInt && vtype != VTFloat {
				return vtype, errors.New("invalid unary operator - on type: " + VTString(vtype))
			}
			op = OpMinus
		case NOT:
			if vtype != VTUnknown && vtype != VTBool {
				return vtype, errors.New("invalid unary operator ! on type: " + VTString(vtype))
			}
			op = OpNot
		default:
			return VTUnknown, errors.New("unknown unary operator")
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
			if i < 0 {
				return VTUnknown, errors.New("unknown variable: " + node.name)
			}
			local = false
		}
		var vtype int
		if local {
			env.addCode(Code{OpCode: OpLoadLVar, Operand: i})
			vtype = env.localvars.vars[i].vtype
		} else {
			env.addCode(Code{OpCode: OpLoadGVar, Operand: i})
			vtype = env.vars.vars[i].vtype
		}
		return vtype, nil
	case BoolExpr:
		if node.value {
			env.addCode(Code{OpCode: OpLoadT})
		} else {
			env.addCode(Code{OpCode: OpLoadF})
		}
		return VTBool, nil
	case IntExpr:
		env.addCode(Code{OpCode: OpLoad, Operand: env.addConst(VInt{node.value})})
		return VTInt, nil
	case FloatExpr:
		env.addCode(Code{OpLoad, env.addConst(VFloat{node.value})})
		return VTFloat, nil
	default:
		return VTUnknown, errors.New(fmt.Sprintf("unknown node type: %+v\n", node))
	}
	return VTUnknown, nil
}
