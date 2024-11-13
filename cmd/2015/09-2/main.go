package main

import (
	"fmt"
	"sort"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	// Looking at the input the graph is both small and fully connected thus
	// we don't have to use a complicated algorithm here.  We can enumerate the
	// possible permutations of the vertices and evaluate the length for each.
	graph := InputToGraph()
	cities := graph.Vertices()

	var distances []int
	puz.EnumeratePermutations(len(cities), func(perm []int) bool {
		distance := 0
		for i := 1; i < len(perm); i++ {
			distance += graph.Edge(cities[perm[i-1]], cities[perm[i]])
		}

		distances = append(distances, distance)
		return false
	})

	fmt.Println(puz.Max(distances...))
}

type Graph struct {
	vertices  puz.Set[string]
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
	sort.Slice(entries, func(i, j int) bool {
		return entries[i] < entries[j]
	})
	return entries
}

func (g *Graph) Edge(from, to string) int {
	var edge int
	if g.distances == nil || g.distances[from] == nil {
		return edge
	}

	return g.distances[from][to]
}

func InputToGraph() Graph {
	var g Graph
	for _, line := range puz.InputToLines(2015, 9) {
		var from, to string
		var distance int

		if _, err := fmt.Sscanf(line, "%s to %s = %d", &from, &to, &distance); err == nil {
			g.AddEdge(from, to, distance)
		}
	}

	return g
}
