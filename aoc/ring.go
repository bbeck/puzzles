package aoc

import (
	"container/list"
	"fmt"
	"strings"
)

// A Ring is a circular buffer of elements along with a reference to a current
// element.
type Ring struct {
	list    *list.List
	current *list.Element
}

func NewRing() *Ring {
	return &Ring{
		list:    list.New(),
		current: nil,
	}
}

func (r *Ring) Current() interface{} {
	return r.current.Value
}

func (r *Ring) InsertAfter(value interface{}) {
	if r.current == nil {
		r.current = r.list.PushBack(value)
	} else {
		r.current = r.list.InsertAfter(value, r.current)
	}
}

func (r *Ring) InsertBefore(value interface{}) {
	if r.current == nil {
		r.current = r.list.PushFront(value)
	} else {
		r.current = r.list.InsertBefore(value, r.current)
	}
}

func (r *Ring) Next() interface{} {
	if r.current == nil {
		r.current = r.list.Front()
	}

	r.current = r.current.Next()
	if r.current == nil {
		r.current = r.list.Front()
	}

	return r.current.Value
}

func (r *Ring) NextN(n int) interface{} {
	var value interface{}
	for i := 0; i < n; i++ {
		value = r.Next()
	}

	return value
}

func (r *Ring) Prev() interface{} {
	if r.current == nil {
		r.current = r.list.Back()
	}

	r.current = r.current.Prev()
	if r.current == nil {
		r.current = r.list.Back()
	}

	return r.current.Value
}

func (r *Ring) PrevN(n int) interface{} {
	var value interface{}
	for i := 0; i < n; i++ {
		value = r.Prev()
	}

	return value
}

// Remove removes the current element from the ring.
func (r *Ring) Remove() interface{} {
	if r.current == nil {
		r.current = r.list.Front()
	}

	remove := r.current
	r.Next()

	return r.list.Remove(remove)
}

func (r *Ring) String() string {
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
