package lib

// A Ring is a circular buffer of elements along with a reference to a current
// element.
type Ring[T any] struct {
	deque Deque[T]
}

// Current returns the element that is currently being reference in the ring.
func (r *Ring[T]) Current() T {
	return r.deque.PeekFront()
}

// InsertAfter inserts a new value in the ring after the current one.  The
// currently referenced element is changed to be the newly inserted one.
func (r *Ring[T]) InsertAfter(value T) {
	r.deque.Rotate(-1)
	r.deque.PushFront(value)
}

// InsertBefore inserts a new value in the ring before the current one.  The
// currently referenced element is changed to be the newly inserted one.
func (r *Ring[T]) InsertBefore(value T) {
	r.deque.PushFront(value)
}

// Next moves the currently referenced element to the next one.
func (r *Ring[T]) Next() T {
	r.deque.Rotate(-1)
	return r.deque.PeekFront()
}

// NextN moves the currently referenced element to the element n steps
// after the currently referenced element.
func (r *Ring[T]) NextN(n int) T {
	r.deque.Rotate(-n)
	return r.deque.PeekFront()
}

// Prev moves the currently referenced element to the previous one.
func (r *Ring[T]) Prev() T {
	r.deque.Rotate(1)
	return r.deque.PeekFront()
}

// PrevN moves the currently referenced element to the element n steps
// before the currently referenced element.
func (r *Ring[T]) PrevN(n int) T {
	r.deque.Rotate(n)
	return r.deque.PeekFront()
}

// Remove removes the current element from the ring.
func (r *Ring[T]) Remove() T {
	return r.deque.PopFront()
}

func (r *Ring[T]) Entries() []T {
	return r.deque.Entries()
}
