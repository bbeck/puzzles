package aoc

// A Stack is a container of elements where the last one added is the first one
// removed.
type Stack struct {
	top *sn
}

type sn struct {
	data  interface{}
	child *sn
}

// NewStack creates a new, empty stack.
func NewStack() *Stack {
	return &Stack{}
}

// Empty determines if there are any elements on the stack.
func (s *Stack) Empty() bool {
	return s.top == nil
}

// Push adds an element to the stack.
func (s *Stack) Push(elem interface{}) {
	s.top = &sn{data: elem, child: s.top}
}

// Peek returns the last pushed element, but does not remove it from the stack.
// If the stack is empty then nil is returned.
func (s *Stack) Peek() interface{} {
	if s.top == nil {
		return nil
	}
	return s.top.data
}

// Pop returns the last pushed element and removes it from the stack.  If the
// stack is empty then nil is returned.
func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}

	elem := s.top.data
	s.top = s.top.child
	return elem
}
