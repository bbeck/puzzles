package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	graph := InputToGraph()
	v1, v2 := SplitGraph(graph, 3)

	c1 := Count(graph, v1)
	c2 := Count(graph, v2)
	fmt.Println(c1 * c2)
}

func SplitGraph(g Graph, mincut int) (string, string) {
	// Since we know how many edges are part of the min cut we know the graph
	// is structured so that there are two sides to it that are separated by
	// mincut edges.  On each side of the graph we know there are at least
	// mincut+1 paths between any two vertices -- if there were fewer then
	// the two vertices would have to be on separate sides of the graph.
	//
	// This algorithm will choose two vertices at random and find a path
	// between them.  It will then remove all edges that are part of the
	// path and then repeat the process.  If we're able to find mincut+1
	// paths then we know that the two vertices chosen at random were on the
	// same side of the graph.  In this case we'll put back all removed edges
	// and try again.
	//
	// Once we know we've found two edges that are on opposite sides of the
	// graph we are done.  Having removed the edges of the path will have
	// broken the graph into its two pieces.

	for {
		v1 := g.ArbitraryVertex()
		v2 := g.ArbitraryVertex()
		isGoal := func(v string) bool { return v == v2 }

		removed := make(map[string]aoc.Set[string])

		var count int
		for count = 0; count < mincut+1; count++ {
			path, ok := aoc.BreadthFirstSearch(v1, g.Children, isGoal)
			if !ok {
				break
			}

			for i := 1; i < len(path); i++ {
				from, to := path[i-1], path[i]
				g.RemoveEdge(from, to)
				g.RemoveEdge(to, from)

				s := removed[from]
				s.Add(to)
				removed[from] = s
			}
		}

		if count == mincut {
			return v1, v2
		}

		// Restore the removed edges
		for from := range removed {
			for to := range removed[from] {
				g.AddEdge(from, to)
				g.AddEdge(to, from)
			}
		}
	}
}

func Count(g Graph, v string) int {
	var count int
	isGoal := func(v string) bool {
		count++
		return false
	}

	aoc.BreadthFirstSearch(v, g.Children, isGoal)
	return count
}

type Graph map[string]aoc.Set[string]

func (g *Graph) ArbitraryVertex() string {
	for v := range *g {
		return v
	}
	return ""
}

func (g *Graph) Children(v string) []string {
	return (*g)[v].Entries()
}

func (g *Graph) AddEdge(from, to string) {
	(*g)[from] = (*g)[from].UnionElems(to)
}

func (g *Graph) RemoveEdge(from, to string) {
	(*g)[from] = (*g)[from].DifferenceElems(to)
}

func InputToGraph() Graph {
	g := make(Graph)
	for _, line := range aoc.InputToLines(2023, 25) {
		line = strings.ReplaceAll(line, ":", "")
		fields := strings.Fields(line)

		lhs := fields[0]
		for _, rhs := range fields[1:] {
			g.AddEdge(lhs, rhs)
			g.AddEdge(rhs, lhs)
		}
	}

	return g
}
