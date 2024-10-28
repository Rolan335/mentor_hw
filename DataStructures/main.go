package main

import (
	"DataStructures/stack"
	"fmt"
)

func main() {
	s := stack.Stack{}
	s.Push("hello")
	res, _ := s.Peek()
	fmt.Println(res)
}
