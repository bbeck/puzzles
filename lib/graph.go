package lib

// EnumerateMaximalCliques calls the fn callback for each maximal clique in a
// graph.
//
// EnumerateMaximalCliques implements the Bron-Kerbosch algorithm with pivoting.
// See: https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func EnumerateMaximalCliques[T comparable](graph map[T]Set[T], fn func([]T)) {
	var zero T

	var BronKerbosch func(R, P, X Set[T])
	BronKerbosch = func(R, P, X Set[T]) {
		// If P and X are both empty, report R as a maximal clique
		if len(P) == 0 && len(X) == 0 {
			fn(R.Entries())
		}

		// Choose a pivot vertex u in P U X
		var u T
		for x := range P {
			u = x
			break
		}
		if u == zero {
			for x := range X {
				u = x
				break
			}
		}

		for v := range P.Difference(graph[u]) {
			BronKerbosch(R.UnionElems(v), P.Intersect(graph[v]), X.Intersect(graph[v]))
			P = P.DifferenceElems(v)
			X = X.UnionElems(v)
		}
	}

	BronKerbosch(nil, SetFrom[T](Keys(graph)...), nil)
}
