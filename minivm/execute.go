package minivm

func Execute(filename string, debug bool) error {
	lexer := new(Lexer)
	if err := lexer.Init(filename); err != nil {
		return err
	}
	Parse(lexer)
	env, err := Codegen(lexer.Result())
	if err != nil {
		return err
	}
	if debug {
		env.Debug()
	}
	env.Execute()
	return nil
}
