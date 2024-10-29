package stack

type Stack struct {
	stack []interface{}
}

func (s *Stack) Push(elem interface{}) {
	s.stack = append(s.stack, elem)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.stack) == 0 {
		return nil, false
	}
	elem := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return elem, true
}

func (s *Stack) Peek() (interface{}, bool) {
	if len(s.stack) == 0 {
		return nil, false
	}
	return s.stack[len(s.stack)-1], true
}

func (s *Stack) Len() int {
	return len(s.stack)
}
