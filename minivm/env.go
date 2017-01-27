package minivm

type Env struct {
	pc       int
	code     []Code
	constant []Value
	stack    *Stack
	vars     *Vars
	breaks   []int
	conts    []int
}

func Codegen(node Node) *Env {
	env := new(Env)
	env.stack = new(Stack)
	env.vars = new(Vars)
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
