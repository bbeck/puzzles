package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	graph := make(map[string]string)
	for _, o := range InputToOrbits() {
		graph[o.Child] = o.Parent
	}

	pYOU, pSAN := PathToRoot(graph, "YOU"), PathToRoot(graph, "SAN")

	// Find the closest common ancestor between YOU and SAN -- ancestor is
	// referenced from the back of the path arrays.
	var ancestor int
	for pYOU[len(pYOU)-ancestor-1] == pSAN[len(pSAN)-ancestor-1] {
		ancestor++
	}

	// The remaining elements are the ones to traverse.
	fmt.Println(len(pYOU) - ancestor + len(pSAN) - ancestor)
}

func PathToRoot(graph map[string]string, n string) []string {
	var path []string
	for n != "" {
		path = append(path, n)
		n = graph[n]
	}
	return path[1:] // don't include the node itself in the path
}

type Orbit struct {
	Parent, Child string
}

func InputToOrbits() []Orbit {
	return aoc.InputLinesTo(2019, 6, func(line string) Orbit {
		var orbit Orbit
		orbit.Parent, orbit.Child, _ = strings.Cut(line, ")")
		return orbit
	})
}
