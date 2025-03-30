package algorithms

import "fmt"

type Stack struct {
	Stack []int
}

func (s *Stack) Add(val int) {
	s.Stack = append(s.Stack, val)
}

func (s *Stack) Pop() {
	if !s.IsEmpty() {
		s.Stack = s.Stack[:len(s.Stack)-1]
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.Stack) == 0
}

func (s *Stack) Peek() int {
	if !s.IsEmpty() {
		return s.Stack[len(s.Stack)-1]
	}
	return -1
}

func StackPrint() {
	var stack Stack
	stack.Add(3)
	stack.Add(4)
	stack.Add(5)

	stack.Pop()
	last := stack.Peek()
	fmt.Println(stack, last)
}
