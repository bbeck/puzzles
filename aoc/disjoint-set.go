package aoc

type DisjointSet struct {
	parent *DisjointSet
	Data   interface{}
	Size   int
}

// NewDisjointSet creates a new disjoint set that contains a single element.
func NewDisjointSet(data interface{}) *DisjointSet {
	set := &DisjointSet{Data: data, Size: 1}
	set.parent = set
	return set
}

// Find returns a representative element that belongs in the same set as the
// provided element.  For any element in the set the same representative element
// will always be returned.
func (e *DisjointSet) Find() *DisjointSet {
	// As we traverse to the root of the tree we'll perform path compression by
	// flattening the tree some by making children point to their grandparent.
	// Over time this will make all children point to the root of the set.
	var elem = e
	for elem.parent != elem {
		elem.parent = elem.parent.parent
		elem = elem.parent
	}

	return elem
}

// Union merges two disjoint subsets into a single set.
func (e *DisjointSet) Union(a *DisjointSet) {
	u := e.Find()
	v := a.Find()

	// Check to see if these two elements are already in the same set.
	if u == v {
		return
	}

	// The set that contains the most children will be the parent
	if u.Size < v.Size {
		u, v = v, u
	}

	v.parent = u
	u.Size += v.Size
}
