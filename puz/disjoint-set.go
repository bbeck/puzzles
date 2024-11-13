package puz

type DisjointSet[T comparable] struct {
	parents map[T]T
	sizes   map[T]int
}

// Add adds an element to the disjoint set.
func (s *DisjointSet[T]) Add(elem T) {
	s.initialize()

	if _, found := s.parents[elem]; !found {
		s.parents[elem] = elem
		s.sizes[elem] = 1
	}
}

// Find returns a representative element that belongs in the same set as the
// provided element.  For any element in the set the same representative
// element will always be returned.  If the element is not in the set then
// the zero value of the generic type will be returned.
func (s *DisjointSet[T]) Find(elem T) (T, bool) {
	s.initialize()

	_, found := s.parents[elem]
	if !found {
		var zero T
		return zero, false
	}

	// Save the path we take to the root of the tree.
	path := []T{elem}
	for {
		parent, found := s.parents[elem]
		if !found || elem == parent {
			break
		}

		elem = parent
		path = append(path, parent)
	}

	// Compress the path we found so that future find operations are near
	// constant time.
	for _, n := range path {
		s.parents[n] = elem
	}

	return elem, true
}

// Union merges two disjoint subsets together into a single set.
func (s *DisjointSet[T]) Union(u, v T) {
	var found bool
	if u, found = s.Find(u); !found {
		return
	}
	if v, found = s.Find(v); !found {
		return
	}

	// Check to see if these two elements are already in the same set.
	if u == v {
		return
	}

	// The set that contains the most children will be the parent
	if s.sizes[u] < s.sizes[v] {
		u, v = v, u
	}

	s.parents[v] = u
	s.sizes[u] += s.sizes[v]
}

// UnionWithAdd will merge two disjoint subsets together into a single set.
// Prior to merging it ensures that each subset is present in the disjoint set.
func (s *DisjointSet[T]) UnionWithAdd(u, v T) {
	s.Add(u)
	s.Add(v)
	s.Union(u, v)
}

// Size returns the size of the subset containing elem.
func (s *DisjointSet[T]) Size(elem T) int {
	if elem, found := s.Find(elem); found {
		return s.sizes[elem]
	}

	return 0
}

func (s *DisjointSet[T]) initialize() {
	if s.parents == nil {
		s.parents = make(map[T]T)
		s.sizes = make(map[T]int)
	}
}
