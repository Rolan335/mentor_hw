package stack

import "errors"

type Stack struct {
	stack []interface{}
}

func (s *Stack) Push(elem interface{}) {
	s.stack = append(s.stack, elem)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.stack) == 0 {
		return nil,errors.New("error. Stack len is 0")
	}
	elem := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return elem, nil
}

func (s Stack) Peek() (interface{}, error) {
	if len(s.stack) == 0 {
		return nil, errors.New("error. Stack len is 0")
	}
	return s.stack[len(s.stack)-1], nil
}

func (s Stack) Len() int{
	return len(s.stack)
}
