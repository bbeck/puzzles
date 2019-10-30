package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	graph := InputToGraph(2018, 7)
	order := TopologicalSort(graph)

	fmt.Printf("order: %s\n", strings.Join(order, ""))
}

type Graph struct {
	vertices map[string]int
	matrix   [][]bool
}

// Compute the topological sort of the graph using Kahn's algorithm -- this
// method will modify the graph.
func TopologicalSort(g *Graph) []string {
	// return the parent node of a given node (i.e. the nodes that the given node
	// has an incoming edge from).
	parents := func(to string) []string {
		tid := g.vertices[to]

		var parents []string
		for from, fid := range g.vertices {
			if g.matrix[fid][tid] {
				parents = append(parents, from)
			}
		}

		return parents
	}

	children := func(from string) []string {
		fid := g.vertices[from]

		var children []string
		for to, tid := range g.vertices {
			if g.matrix[fid][tid] {
				children = append(children, to)
			}
		}

		return children
	}

	var L []string // the topological sort output

	var S []string // set of nodes that have no incoming edge
	for name := range g.vertices {
		if len(parents(name)) == 0 {
			S = append(S, name)
		}
	}

	for len(S) > 0 {
		// ensure S is sorted
		sort.Strings(S)

		// remove a node n from S
		n := S[0]
		S = S[1:]

		// add n to tail of L
		L = append(L, n)

		// for each node m with an edge e from n to m
		for _, m := range children(n) {
			// remove edge e from the graph
			g.matrix[g.vertices[n]][g.vertices[m]] = false

			// if m has no other incoming edges then
			if len(parents(m)) == 0 {
				// insert m into S
				S = append(S, m)
			}
		}
	}

	return L
}

func InputToGraph(year, day int) *Graph {
	assigned := make(map[string]bool)  // each vertex that has been assigned an id
	vertices := make(map[string]int)   // the assigned id for each vertex
	edges := make(map[string][]string) // the mapping of edges from vertex to vertex

	for _, line := range aoc.InputToLines(year, day) {
		var from, to string
		if _, err := fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &from, &to); err != nil {
			log.Fatalf("unable to parse line: %s", line)
		}

		if !assigned[from] {
			vertices[from] = len(vertices)
			assigned[from] = true
		}
		if !assigned[to] {
			vertices[to] = len(vertices)
			assigned[to] = true
		}

		edges[from] = append(edges[from], to)
	}

	matrix := make([][]bool, len(vertices))
	for name, id := range vertices {
		matrix[id] = make([]bool, len(vertices))
		for _, child := range edges[name] {
			matrix[id][vertices[child]] = true
		}
	}

	return &Graph{
		vertices: vertices,
		matrix:   matrix,
	}
}
