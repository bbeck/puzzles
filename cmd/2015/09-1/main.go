package main

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	// Looking at the input the graph is both small and fully connected so we
	// don't have to use a complicated algorithm here.  We can enumerate the
	// possible permutations of the vertices and evaluate the length for each.
	vertices, weights := InputToGraph(2015, 9)

	best := math.MaxInt64
	aoc.EnumeratePermutations(len(vertices), func(path []int) {
		cost := Cost(path, weights)
		if cost < best {
			best = cost
		}
	})

	fmt.Printf("best: %d\n", best)
}

func InputToGraph(year, day int) ([]string, [][]int) {
	type entry struct {
		from, to string
		distance int
	}

	// Parse the lines into entries so that we can make multiple passes over them.
	entries := make([]entry, 0)
	for _, line := range aoc.InputToLines(year, day) {
		var from, to string
		var distance int
		if _, err := fmt.Sscanf(line, "%s to %s = %d", &from, &to, &distance); err != nil {
			log.Fatalf("unable to parse line: %s", line)
		}

		entries = append(entries, entry{from, to, distance})
	}

	// Determine the unique set of vertices and the index for each one
	indices := make(map[string]int)
	for _, entry := range entries {
		if fromIndex, ok := indices[entry.from]; !ok {
			fromIndex = len(indices)
			indices[entry.from] = fromIndex
		}

		if toIndex, ok := indices[entry.to]; !ok {
			toIndex = len(indices)
			indices[entry.to] = toIndex
		}
	}

	// Determine the vertices ordered by their id.
	vertices := make([]string, 0)
	for vertex := range indices {
		vertices = append(vertices, vertex)
	}
	sort.Slice(vertices, func(i, j int) bool {
		return indices[vertices[i]] < indices[vertices[j]]
	})

	// Now build the square matrix of weights
	weights := make([][]int, len(indices))
	for i := 0; i < len(indices); i++ {
		weights[i] = make([]int, len(indices))
	}
	for _, entry := range entries {
		// The graph is bidirectional.
		weights[indices[entry.from]][indices[entry.to]] = entry.distance
		weights[indices[entry.to]][indices[entry.from]] = entry.distance
	}

	return vertices, weights
}

func Cost(path []int, weights [][]int) int {
	var cost int

	from := path[0]
	for i := 1; i < len(path); i++ {
		to := path[i]
		cost += weights[from][to]
		from = to
	}

	return cost
}
