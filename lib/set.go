package lib

import (
	"fmt"
	"math/bits"
	"sort"
	"strings"
)

var present = struct{}{}

type Set[T comparable] map[T]struct{}

// SetFrom creates a new set with the arguments as elements.
func SetFrom[T comparable](elems ...T) Set[T] {
	var s Set[T]
	s.Add(elems...)
	return s
}

// Add adds a new element to the set.  Returns true if the set was modified.
func (s *Set[T]) Add(elems ...T) bool {
	if *s == nil {
		*s = make(Set[T])
	}

	var changed bool
	for _, elem := range elems {
		if _, found := (*s)[elem]; !found {
			(*s)[elem] = present
			changed = true
		}
	}

	return changed
}

// Remove removes the provided element from the set.  Returns true if the set was modified.
func (s *Set[T]) Remove(elems ...T) bool {
	if *s == nil {
		*s = make(Set[T])
		return false
	}

	var changed bool
	for _, elem := range elems {
		if _, found := (*s)[elem]; found {
			delete(*s, elem)
			changed = true
		}
	}

	return changed
}

// Contains returns true if the set contains the provided element.
func (s Set[T]) Contains(elem T) bool {
	_, found := s[elem]
	return found
}

// Entries returns the entries in the set.
func (s Set[T]) Entries() []T {
	entries := make([]T, 0, len(s))
	for entry := range s {
		entries = append(entries, entry)
	}

	return entries
}

// String returns a human-readable string of the set.
func (s Set[T]) String() string {
	var keys []string
	for key := range s {
		keys = append(keys, fmt.Sprintf("%v", key))
	}
	sort.Strings(keys)

	var sb strings.Builder
	sb.WriteRune('{')
	sb.WriteString(strings.Join(keys, ", "))
	sb.WriteRune('}')
	return sb.String()
}

// Union returns a new set that contains all the elements of the current set
// as well as all the elements from the provided set.
func (s Set[T]) Union(other Set[T]) Set[T] {
	var union Set[T]
	for elem := range s {
		union.Add(elem)
	}
	for elem := range other {
		union.Add(elem)
	}

	return union
}

// UnionElems returns a new set that contains all the elements of the current
// set as well as all the elements provided as arguments.
func (s Set[T]) UnionElems(elems ...T) Set[T] {
	other := SetFrom(elems...)
	return s.Union(other)
}

// Intersect returns a new set that contains only the elements that are present
// in the current set as well as the provided set.
func (s Set[T]) Intersect(other Set[T]) Set[T] {
	var intersection Set[T]
	for elem := range s {
		if other.Contains(elem) {
			intersection.Add(elem)
		}
	}

	return intersection
}

// IntersectElems returns a new set that contains only the elements that are
// present in the current set as ewll as provided as arguments.
func (s Set[T]) IntersectElems(elems ...T) Set[T] {
	other := SetFrom(elems...)
	return s.Intersect(other)
}

// Difference returns a new set that contains only the elements present in the
// current set that are not in the provided set.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	var difference Set[T]
	for elem := range s {
		if !other.Contains(elem) {
			difference.Add(elem)
		}
	}

	return difference
}

// DifferenceElems returns a new set that contains only the elements present in the
// current set that are not in the provided arguments.
func (s Set[T]) DifferenceElems(elems ...T) Set[T] {
	other := SetFrom(elems...)
	return s.Difference(other)
}

// BitSet is a representation of a set that can hold up to 64 elements.  BitSet
// operates differently from a generic set in that its elements must be
// integers on the range of 0 to 63 (inclusive).  This is because a BitSet is
// represented by a 64-bit integer.  Additionally, as a consequence of this
// representation, BitSet instances are always passed by value.
type BitSet uint64

// Add returns a new BitSet that includes the specified element.
func (bs BitSet) Add(elem int) BitSet {
	return bs | (1 << elem)
}

// Remove returns a new BitSet that doesn't include the specified element.
func (bs BitSet) Remove(elem int) BitSet {
	return bs & ^(1 << elem)
}

// Contains returns true if the current BitSet contains the specified element.
func (bs BitSet) Contains(elem int) bool {
	return bs&(1<<elem) != 0
}

// Size returns the number of elements in the BitSet.
func (bs BitSet) Size() int {
	return bits.OnesCount64(uint64(bs))
}
