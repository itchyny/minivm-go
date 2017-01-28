package minivm

import "math"

func optimize(node Node) Node {
	switch node := node.(type) {
	case Function:
		return Function{name: node.name, args: node.args, stmts: optimize(node.stmts)}
	case ReturnStmt:
		return ReturnStmt{expr: optimize(node.expr)}
	case Statements:
		var stmts []Node
		for _, stmt := range node.stmts {
			stmts = append(stmts, optimize(stmt))
		}
		return Statements{stmts: stmts}
	case IfStmt:
		return IfStmt{expr: optimize(node.expr), stmts: optimize(node.stmts), elsestmts: optimize(node.elsestmts)}
	case WhileStmt:
		return WhileStmt{expr: optimize(node.expr), stmts: optimize(node.stmts)}
	case LetStmt:
		return LetStmt{ident: node.ident, expr: optimize(node.expr)}
	case PrintStmt:
		return PrintStmt{expr: optimize(node.expr)}
	case CallExpr:
		var exprs []Node
		for _, expr := range node.exprs {
			exprs = append(exprs, optimize(expr))
		}
		return CallExpr{name: node.name, exprs: exprs}
	case BinOpExpr:
		return node.optimize()
	case UnaryOpExpr:
		return node.optimize()
	default:
		return node
	}
}

func (expr BinOpExpr) optimize() Node {
	var node Node
	left := optimize(expr.left)
	right := optimize(expr.right)
	node = BinOpExpr{op: expr.op, left: left, right: right}
	switch left := left.(type) {
	case BoolExpr:
		if expr.op == AND {
			if left.value {
				node = right
			} else {
				node = BoolExpr{value: false}
			}
		} else if expr.op == OR {
			if left.value {
				node = BoolExpr{value: true}
			} else {
				node = right
			}
		}
	case IntExpr:
		if math.MinInt32 <= left.value && left.value <= math.MaxInt32 &&
			(expr.op == PLUS || expr.op == TIMES || expr.op == EQEQ || expr.op == NEQ) {
			node = BinOpExprI{op: expr.op, left: right, right: int(left.value)}
		}
	default:
		switch right := right.(type) {
		case IntExpr:
			if math.MinInt32 <= right.value && right.value <= math.MaxInt32 {
				node = BinOpExprI{op: expr.op, left: left, right: int(right.value)}
			}
		}
	}
	return node
}

func (expr UnaryOpExpr) optimize() Node {
	var node Node
	node = expr
	switch e := optimize(expr.expr).(type) {
	case BoolExpr:
		if expr.op == NOT {
			node = BoolExpr{value: !e.value}
		}
	case IntExpr:
		if expr.op == UPLUS {
			node = IntExpr{value: e.value}
		} else if expr.op == UMINUS {
			node = IntExpr{value: -e.value}
		}
	case FloatExpr:
		if expr.op == UPLUS {
			node = FloatExpr{value: e.value}
		} else if expr.op == UMINUS {
			node = FloatExpr{value: -e.value}
		}
	}
	return node
}
