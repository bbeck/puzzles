package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"sort"
)

func main() {
	durations := make(map[string]int)
	for i, c := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		durations[string(c)] = 60 + i + 1
	}

	graph := InputToGraph()
	workers := make([]Worker, 5) // the duration remaining on each worker's task

	tm := -1 // first iteration is setting things up for the first tick
	for {
		idle := true
		for i := range workers {
			workers[i].Remaining--

			if workers[i].Remaining <= 0 {
				graph.Remove(workers[i].Vertex)
				workers[i].Vertex = graph.Choose()
				workers[i].Remaining = durations[workers[i].Vertex]
			}

			idle = idle && workers[i].Remaining == 0
		}

		if graph.IsEmpty() && idle {
			break
		}

		tm++
	}

	fmt.Println(tm)
}

type Worker struct {
	Vertex    string
	Remaining int
}

type Graph struct {
	Vertices puz.Set[string]
	Parents  map[string][]string
}

func (g *Graph) IsEmpty() bool {
	for _, parents := range g.Parents {
		if len(parents) > 0 {
			return false
		}
	}
	return true
}

func (g *Graph) AddEdge(child, parent string) {
	if g.Parents == nil {
		g.Parents = make(map[string][]string)
	}

	g.Vertices.Add(child, parent)
	g.Parents[child] = append(g.Parents[child], parent)
}

func (g *Graph) Choose() string {
	var candidates []string
	for child := range g.Vertices {
		if len(g.Parents[child]) == 0 {
			candidates = append(candidates, child)
		}
	}
	sort.Strings(candidates)

	if candidates == nil {
		return ""
	}

	choice := candidates[0]
	g.Vertices.Remove(choice)
	return choice
}

func (g *Graph) Remove(vertex string) {
	for child := range g.Parents {
		g.Parents[child] = Remove(g.Parents[child], vertex)
	}
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
	for _, line := range puz.InputToLines() {
		var parent, child string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &parent, &child)

		graph.AddEdge(child, parent)
	}

	return graph
}
