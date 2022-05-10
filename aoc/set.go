package aoc

import (
	"fmt"
	"sort"
	"strings"
)

var present = struct{}{}

type Set[T comparable] map[T]struct{}

// SingletonSet creates a new set with a single element.
func SingletonSet[T comparable](elem T) Set[T] {
	return map[T]struct{}{elem: present}
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
