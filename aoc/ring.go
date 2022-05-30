package aoc

import (
	"container/list"
	"fmt"
	"log"
	"strings"
)

// A Ring is a circular buffer of elements along with a reference to a current
// element.
//
// Additionally, it stores a lookup table to make it possible to jump to a
// specific entry in the ring.  This lookup only works when the values in the
// ring are distinct from one another.
type Ring[T comparable] struct {
	list    *list.List
	current *list.Element
	lookup  map[T]*list.Element
}

func (r *Ring[T]) initialize() {
	if r.list == nil {
		r.list = list.New()
	}

	if r.lookup == nil {
		r.lookup = make(map[T]*list.Element)
	}
}

// Current returns the element that is currently being reference in the ring.
func (r *Ring[T]) Current() T {
	return r.current.Value.(T)
}

// JumpTo changes the current element in the ring to the specified one.  If the
// specified value isn't found in the ring then the process will exit.
func (r *Ring[T]) JumpTo(value T) {
	r.initialize()

	elem, found := r.lookup[value]
	if !found {
		log.Fatalf("unable to find value %v in lookup table", value)
		return
	}

	r.current = elem
}

// InsertAfter inserts a new value in the ring after the current one.  The
// currently referenced element is changed to be the newly inserted one.
func (r *Ring[T]) InsertAfter(value T) {
	r.initialize()

	if r.current == nil {
		r.current = r.list.PushBack(value)
	} else {
		r.current = r.list.InsertAfter(value, r.current)
	}
	r.lookup[value] = r.current
}

// InsertBefore inserts a new value in the ring before the current one.  The
// currently referenced element is changed to be the newly inserted one.
func (r *Ring[T]) InsertBefore(value T) {
	r.initialize()

	if r.current == nil {
		r.current = r.list.PushFront(value)
	} else {
		r.current = r.list.InsertBefore(value, r.current)
	}
	r.lookup[value] = r.current
}

// Next moves the currently referenced element to the next one.
func (r *Ring[T]) Next() T {
	r.initialize()

	if r.current == nil {
		r.current = r.list.Front()
	}

	r.current = r.current.Next()
	if r.current == nil {
		r.current = r.list.Front()
	}

	return r.current.Value.(T)
}

// NextN moves the currently referenced element to the element n steps
// after the currently referenced element.
func (r *Ring[T]) NextN(n int) T {
	var value T
	for i := 0; i < n; i++ {
		value = r.Next()
	}

	return value
}

// Prev moves the currently referenced element to the previous one.
func (r *Ring[T]) Prev() T {
	r.initialize()

	if r.current == nil {
		r.current = r.list.Back()
	}

	r.current = r.current.Prev()
	if r.current == nil {
		r.current = r.list.Back()
	}

	return r.current.Value.(T)
}

// PrevN moves the currently referenced element to the element n steps
// before the currently referenced element.
func (r *Ring[T]) PrevN(n int) T {
	var value T
	for i := 0; i < n; i++ {
		value = r.Prev()
	}

	return value
}

// Remove removes the current element from the ring.
func (r *Ring[T]) Remove() T {
	r.initialize()

	if r.current == nil {
		r.current = r.list.Front()
	}

	remove := r.current
	r.Next()

	value := r.list.Remove(remove).(T)
	delete(r.lookup, value)
	return value
}

// String converts the ring to a string.
func (r *Ring[T]) String() string {
	r.initialize()

	var builder strings.Builder
	for elem := r.list.Front(); elem != nil; elem = elem.Next() {
		if elem == r.current {
			builder.WriteRune('<')
		}

		switch v := elem.Value.(type) {
		case fmt.Stringer:
			builder.WriteString(v.String())
		default:
			builder.WriteString(fmt.Sprintf("%+v", v))
		}

		if elem == r.current {
			builder.WriteRune('>')
		}

		if elem.Next() != nil {
			builder.WriteRune(' ')
		}
	}

	return builder.String()
}
