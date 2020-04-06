package stack

import (
	"fmt"
	"hw5/list"
)

// Stack interface
type Stack interface {
	IsEmpty() bool
	Peek() (interface{}, error)
	Pop() (interface{}, error)
	Push(item interface{})
}

// stack implementation
type stack struct {
	top *list.Node
}

// New creates new stack
func New() Stack {
	return &stack{}
}

// IsEmpty returns true if the stack is empty
func (s *stack) IsEmpty() bool {
	return s.top == nil
}

// Peek returns top element on the stack
func (s *stack) Peek() (interface{}, error) {
	if s.top == nil {
		return nil, fmt.Errorf("The stack is empty")
	}
	return s.top.Data, nil
}

// Pop removes top element from the stack
func (s *stack) Pop() (interface{}, error) {
	if s.top == nil {
		return nil, fmt.Errorf("The stack is empty")
	}
	n := s.top
	s.top = n.Next
	n.Next = nil
	return n.Data, nil
}

// Push adds new item on top of the stack
func (s *stack) Push(item interface{}) {
	n := &list.Node{Data: item}
	n.Next = s.top
	s.top = n
}
