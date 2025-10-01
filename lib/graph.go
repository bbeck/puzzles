package lib

import "sort"

type WeightedEdge[T any] struct {
	From   T
	To     T
	Weight int
}

// MinimumSpanningTree computes the minimum spanning tree of a graph using
// Kruskal's algorithm.  It returns the total cost of the tree and the edges
// that comprise the tree.
func MinimumSpanningTree[T comparable](vertices []T, children ChildrenFunc[T], cost CostFunc[T]) (int, []WeightedEdge[T]) {
	var ds DisjointSet[T]
	for i := range vertices {
		ds.Add(vertices[i])
	}

	// Determine the edges and order them by weight (ascending).
	var edges []WeightedEdge[T]
	for _, from := range vertices {
		for _, to := range children(from) {
			edges = append(edges, WeightedEdge[T]{From: from, To: to, Weight: cost(from, to)})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	var total int
	var tree []WeightedEdge[T]
	for _, e := range edges {
		u, _ := ds.Find(e.From)
		v, _ := ds.Find(e.To)
		if u != v {
			ds.Union(e.From, e.To)
			total += e.Weight
			tree = append(tree, e)
		}
	}

	return total, tree
}

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
