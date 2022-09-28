package aoc

// A Stack is a container of elements where the last one added is the first one
// removed.
type Stack[T any] struct {
	deque Deque[T]
}

// Empty determines if there are any elements on the stack.
func (s *Stack[T]) Empty() bool {
	return s.deque.Empty()
}

// Len returns the number of elements on the stack.
func (s *Stack[T]) Len() int {
	return s.deque.Len()
}

// Push adds an element to the stack.
func (s *Stack[T]) Push(elem T) {
	s.deque.PushBack(elem)
}

// Peek returns the last pushed element, but does not remove it from the stack.
// If the stack is empty then the zero value is returned.
func (s *Stack[T]) Peek() T {
	return s.deque.PeekBack()
}

// Pop returns the last pushed element and removes it from the stack.  If the
// stack is empty then the zero value is returned.
func (s *Stack[T]) Pop() T {
	return s.deque.PopBack()
}
