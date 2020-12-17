package aoc

import (
	"fmt"
	"strings"
)

var exists = struct{}{}

type Set struct {
	entries map[interface{}]struct{}
}

// NewSet creates a new, empty set.
func NewSet() Set {
	return Set{
		entries: make(map[interface{}]struct{}),
	}
}

// Add adds a new element to the set.
func (s *Set) Add(elem interface{}) {
	s.entries[elem] = exists
}

// Remove removes the provided element from the set.
func (s *Set) Remove(elem interface{}) {
	delete(s.entries, elem)
}

// Contains returns true if the set contains the provided element.
func (s Set) Contains(elem interface{}) bool {
	_, found := s.entries[elem]
	return found
}

// Size returns the number of entries in the set.
func (s Set) Size() int {
	return len(s.entries)
}

// Entries returns the entries in the set.
func (s Set) Entries() []interface{} {
	var entries []interface{}
	for entry := range s.entries {
		entries = append(entries, entry)
	}

	return entries
}

// String returns a human readable string of the set.
func (s Set) String() string {
	var keys []string
	for key := range s.entries {
		keys = append(keys, fmt.Sprintf("%s", key))
	}

	var sb strings.Builder
	sb.WriteRune('{')
	sb.WriteString(strings.Join(keys, ", "))
	sb.WriteRune('}')
	return sb.String()
}

// Union returns a new set that contains all of the elements of the current set
// as well as all of the elements from the provided set.
func (s Set) Union(other Set) Set {
	union := NewSet()
	for elem := range s.entries {
		union.Add(elem)
	}
	for elem := range other.entries {
		union.Add(elem)
	}

	return union
}

// Intersect returns a new set that contains only the elements that are present
// in the current set as well as the provided set.
func (s Set) Intersect(other Set) Set {
	intersection := NewSet()
	for elem := range s.entries {
		if other.Contains(elem) {
			intersection.Add(elem)
		}
	}

	return intersection
}

// Difference returns a new set that contains only the elements present in the
// current set that are not in the provided set.
func (s Set) Difference(other Set) Set {
	difference := NewSet()
	for elem := range s.entries {
		if !other.Contains(elem) {
			difference.Add(elem)
		}
	}

	return difference
}
