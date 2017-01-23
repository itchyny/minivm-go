package minivm

type Stack struct {
	stack []Value
}

func (s *Stack) Empty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Push(value Value) {
	s.stack = append(s.stack, value)
}

func (s *Stack) Pop() Value {
	value := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return value
}

func (s *Stack) Dup() {
	s.stack = append(s.stack, s.stack[len(s.stack)-1])
}
