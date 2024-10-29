package main

import (
	"datastruct/stack"
	"fmt"
)

func main() {
	s := stack.Stack{}
	s.Push("hello")
	res, _ := s.Peek()
	fmt.Println(res)
}
