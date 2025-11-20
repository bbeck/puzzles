package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	graph := BuildGraph(
		in.ToGrid2D(func(_, _ int, s string) int { return ParseInt(s) }),
	)

	var exploded Set[Point2D]
	for range 3 {
		exploded = FindBest(graph, exploded)
	}
	fmt.Println(len(exploded))
}

func FindBest(graph map[Point2D]*Node, exploded Set[Point2D]) Set[Point2D] {
	var bestExploded Set[Point2D]

	for _, start := range graph {
		if exploded.Contains(start.ID) {
			continue
		}

		var seen Set[Point2D]
		var queue = []Point2D{start.ID}
		for len(queue) > 0 {
			var p Point2D
			p, queue = queue[0], queue[1:]

			if seen.Add(p) {
				queue = append(queue, graph[p].NeighborIDs...)
			}
		}

		var e Set[Point2D]
		for p := range seen {
			e.Add(graph[p].Members...)
		}
		e = e.Difference(exploded)

		if len(e) > len(bestExploded) {
			bestExploded = e
		}
	}

	return bestExploded.Union(exploded)
}

type Node struct {
	ID          Point2D
	Members     []Point2D
	NeighborIDs []Point2D
}

func BuildGraph(grid Grid2D[int]) map[Point2D]*Node {
	// Start by building a disjoint set to compute connected components.
	var ds DisjointSet[Point2D]
	grid.ForEachPoint(func(p Point2D, pv int) {
		ds.Add(p)

		grid.ForEachOrthogonalNeighborPoint(p, func(q Point2D, qv int) {
			if pv == qv {
				ds.UnionWithAdd(p, q)
			}
		})
	})

	// Next, collect the connected components into graph nodes.
	var nodes = make(map[Point2D]*Node)
	grid.ForEachPoint(func(p Point2D, pv int) {
		root, _ := ds.Find(p)
		if node, ok := nodes[root]; ok {
			node.Members = append(node.Members, p)
			return
		}

		nodes[root] = &Node{ID: root, Members: []Point2D{p}}
	})

	// Finally, determine the neighbor relationships between the graph nodes.
	// A node can only be neighbors with nodes that are less than or equal to
	// its own value.
	grid.ForEachPoint(func(p Point2D, pv int) {
		pRoot, _ := ds.Find(p)

		grid.ForEachOrthogonalNeighborPoint(p, func(q Point2D, qv int) {
			qRoot, _ := ds.Find(q)

			if pRoot != qRoot && pv >= qv {
				nodes[pRoot].NeighborIDs = append(nodes[pRoot].NeighborIDs, qRoot)
			}
		})
	})

	return nodes
}
