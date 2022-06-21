package aoc

// A Stack is a container of elements where the last one added is the first one
// removed.
type Stack[T any] []T

// Empty determines if there are any elements on the stack.
func (s *Stack[T]) Empty() bool {
	return len(*s) == 0
}

// Push adds an element to the stack.
func (s *Stack[T]) Push(elem T) {
	*s = append(*s, elem)
}

// Peek returns the last pushed element, but does not remove it from the stack.
// If the stack is empty then the zero value is returned.
func (s *Stack[T]) Peek() T {
	n := len(*s)
	if n == 0 {
		var zero T
		return zero
	}

	return (*s)[n-1]
}

// Pop returns the last pushed element and removes it from the stack.  If the
// stack is empty then the zero value is returned.
func (s *Stack[T]) Pop() T {
	n := len(*s)
	if n == 0 {
		var zero T
		return zero
	}

	elem := (*s)[n-1]
	*s = (*s)[:n-1]
	return elem
}
