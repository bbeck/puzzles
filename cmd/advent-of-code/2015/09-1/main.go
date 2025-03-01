package main

import (
	"fmt"

	"slices"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	// Looking at the input the graph is both small and fully connected thus
	// we don't have to use a complicated algorithm here.  We can enumerate the
	// possible permutations of the vertices and evaluate the length for each.
	graph := InputToGraph()
	cities := graph.Vertices()

	var distances []int
	EnumeratePermutations(len(cities), func(perm []int) bool {
		distance := 0
		for i := 1; i < len(perm); i++ {
			distance += graph.Edge(cities[perm[i-1]], cities[perm[i]])
		}

		distances = append(distances, distance)
		return false
	})

	fmt.Println(Min(distances...))
}

type Graph struct {
	vertices  Set[string]
	distances map[string]map[string]int
}

func (g *Graph) AddEdge(from, to string, distance int) {
	g.vertices.Add(from, to)

	if g.distances == nil {
		g.distances = make(map[string]map[string]int)
	}

	if g.distances[from] == nil {
		g.distances[from] = make(map[string]int)
	}
	g.distances[from][to] = distance

	if g.distances[to] == nil {
		g.distances[to] = make(map[string]int)
	}
	g.distances[to][from] = distance
}

func (g *Graph) Vertices() []string {
	entries := g.vertices.Entries()
	slices.Sort(entries)
	return entries
}

func (g *Graph) Edge(from, to string) int {
	if g.distances == nil || g.distances[from] == nil {
		return 0
	}

	return g.distances[from][to]
}

func InputToGraph() Graph {
	var g Graph
	for in.HasNext() {
		var from, to string
		var distance int
		in.Scanf("%s to %s = %d", &from, &to, &distance)
		g.AddEdge(from, to, distance)
	}

	return g
}
