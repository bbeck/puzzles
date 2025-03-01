package main

import (
	"fmt"
	"math"
	"slices"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	graph := InputToGraph()
	people := graph.Vertices()

	// Insert myself as a person as well
	for _, person := range people {
		graph.AddEdge("Me", person, 0)
		graph.AddEdge(person, "Me", 0)
	}
	people = graph.Vertices()

	best := math.MinInt
	EnumeratePermutations(len(people), func(perm []int) bool {
		best = Max(best, Happiness(graph, perm))
		return false
	})

	fmt.Println(best)
}

func Happiness(graph Graph, perm []int) int {
	people := graph.Vertices()
	N := len(people)

	var happiness int
	for i := range perm {
		person := people[perm[i]]
		left := people[perm[(i-1+N)%N]]
		right := people[perm[(i+1)%N]]

		happiness += graph.Edge(person, left)
		happiness += graph.Edge(person, right)
	}

	return happiness
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
		var gainlose string
		var amount int

		in.Scanf("%s would %s %d happiness units by sitting next to %s.", &from, &gainlose, &amount, &to)
		if gainlose == "lose" {
			amount = -amount
		}
		g.AddEdge(from, to, amount)
	}

	return g
}
