package minivm

import "errors"

type Vars struct {
	vars []Var
}

func (vars *Vars) lookup(name string) int {
	for i, v := range vars.vars {
		if v.name == name {
			return i
		}
	}
	return -1
}

func (vars *Vars) set(name string, vtype int) {
	if vars.lookup(name) < 0 {
		vars.vars = append(vars.vars, Var{name: name, vtype: vtype})
	}
}

func (vars *Vars) alloc(node Node) {
	switch node := node.(type) {
	case Function:
		vars.set(node.name, VTFunc)
	case Statements:
		for _, stmt := range node.stmts {
			vars.alloc(stmt)
		}
	case IfStmt:
		vars.alloc(node.stmts)
		vars.alloc(node.elsestmts)
	case WhileStmt:
		vars.alloc(node.stmts)
	case LetStmt:
		vars.set(node.ident, VTUnknown)
	}
}

func (vars *Vars) allocLocal(node Function) error {
	for _, arg := range node.args {
		if vars.lookup(arg) >= 0 {
			return errors.New("duplicated argument name: " + arg)
		}
		vars.set(arg, VTUnknown)
	}
	vars.alloc(node.stmts)
	return nil
}

type Var struct {
	name  string
	vtype int
	value Value
}
