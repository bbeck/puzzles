package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"sort"
	"strings"
)

func main() {
	graph := InputToGraph()
	ordering := graph.TopologicalSort()
	fmt.Println(strings.Join(ordering, ""))
}

type Graph struct {
	Vertices lib.Set[string]
	Parents  map[string][]string
}

func (g *Graph) AddEdge(child, parent string) {
	if g.Parents == nil {
		g.Parents = make(map[string][]string)
	}

	g.Vertices.Add(child, parent)
	g.Parents[child] = append(g.Parents[child], parent)
}

func (g Graph) TopologicalSort() []string {
	var vertices []string
	for len(g.Vertices) > 0 {
		choice := g.Choose()
		vertices = append(vertices, choice)

		g.Vertices.Remove(choice)
		for child := range g.Parents {
			g.Parents[child] = Remove(g.Parents[child], choice)
		}
	}

	return vertices
}

func (g *Graph) Choose() string {
	var candidates []string
	for child := range g.Vertices {
		if len(g.Parents[child]) == 0 {
			candidates = append(candidates, child)
		}
	}

	sort.Strings(candidates)
	return candidates[0]
}

func Remove[T comparable](s []T, elem T) []T {
	var i int
	for i < len(s) {
		if s[i] != elem {
			i++
			continue
		}

		s = append(s[:i], s[i+1:]...)
	}

	return s
}

func InputToGraph() Graph {
	var graph Graph
	for _, line := range lib.InputToLines() {
		var parent, child string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &parent, &child)

		graph.AddEdge(child, parent)
	}

	return graph
}
