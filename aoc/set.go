package aoc

type Set map[string]struct{}

var exists = struct{}{}

// NewSet creates a new, empty set.
func NewSet() *Set {
	return &Set{}
}

// Add adds a new element to the set.
func (s *Set) Add(elem string) {
	(*s)[elem] = exists
}

// Union returns a new set that contains all of the elements of the current set
// as well as all of the elements from the provided set.
func (s *Set) Union(other *Set) *Set {
	union := NewSet()
	for elem := range *s {
		(*union)[elem] = exists
	}
	for elem := range *other {
		(*union)[elem] = exists
	}

	return union
}

// Intersect returns a new set that contains only the elements that are present
// in the current set as well as the provided set.
func (s *Set) Intersect(other *Set) *Set {
	intersection := NewSet()
	for elem := range *s {
		if _, contains := (*other)[elem]; contains {
			(*intersection)[elem] = exists
		}
	}
	for elem := range *other {
		if _, contains := (*s)[elem]; contains {
			(*intersection)[elem] = exists
		}
	}

	return intersection
}
