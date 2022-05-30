package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
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
	aoc.EnumeratePermutations(len(people), func(perm []int) bool {
		best = aoc.Max(best, Happiness(graph, perm))
		return false
	})

	fmt.Println(best)
}

func Happiness(graph Graph, perm []int) int {
	people := graph.Vertices()
	N := len(people)

	var happiness int
	for i := 0; i < len(perm); i++ {
		person := people[perm[i]]
		left := people[perm[(i-1+N)%N]]
		right := people[perm[(i+1)%N]]

		happiness += graph.Edge(person, left)
		happiness += graph.Edge(person, right)
	}

	return happiness
}

type Graph struct {
	vertices  aoc.Set[string]
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
	for _, line := range aoc.InputToLines(2015, 13) {
		// Transform the line to remove filler words to make it easier to parse by Sscanf.
		line = strings.ReplaceAll(line, " would ", " ")
		line = strings.ReplaceAll(line, " happiness units by sitting next to ", " ")
		line = strings.ReplaceAll(line, ".", "")

		var from, to string
		var gainlose string
		var amount int

		if _, err := fmt.Sscanf(line, "%s %s %d %s", &from, &gainlose, &amount, &to); err == nil {
			if gainlose == "lose" {
				amount = -amount
			}
			g.AddEdge(from, to, amount)
		}
	}

	return g
}
