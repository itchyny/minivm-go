package minivm

type Env struct {
	pc        int
	code      []Code
	constant  []Value
	stack     *Stack
	vars      *Vars
	localvars *Vars
	diff      int
	diffs     []int
	returns   []int
	breaks    []int
	conts     []int
}

func Codegen(node Node) (*Env, error) {
	env := new(Env)
	env.stack = new(Stack)
	env.vars = new(Vars)
	env.vars.alloc(node)
	env.diffs = []int{len(env.vars.vars)}
	if _, err := env.codegen(optimize(node)); err != nil {
		return nil, err
	}
	return env, nil
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
