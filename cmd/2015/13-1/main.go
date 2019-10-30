package main

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	people, weights := InputToGraph(2015, 13)

	best := math.MinInt64
	aoc.EnumeratePermutations(len(people), func(perm []int) {
		cost := Cost(perm, weights)
		if cost > best {
			best = cost
		}
	})

	fmt.Printf("best: %d\n", best)
}

func InputToGraph(year, day int) ([]string, [][]int) {
	type entry struct {
		p1, p2    string
		happiness int
	}

	// Parse the lines into entries so that we can make multiple passes over them.
	entries := make([]entry, 0)
	for _, line := range aoc.InputToLines(year, day) {
		var p1, p2, dir string
		var happiness int
		if _, err := fmt.Sscanf(line[:len(line)-1], "%s would %s %d happiness units by sitting next to %s", &p1, &dir, &happiness, &p2); err != nil {
			log.Fatalf("unable to parse line: %s", line)
		}

		if dir == "lose" {
			happiness = -happiness
		}

		entries = append(entries, entry{p1, p2, happiness})
	}

	// Determine the unique set of people and the index for each one
	indices := make(map[string]int)
	for _, entry := range entries {
		if p1Index, ok := indices[entry.p1]; !ok {
			p1Index = len(indices)
			indices[entry.p1] = p1Index
		}

		if p2Index, ok := indices[entry.p2]; !ok {
			p2Index = len(indices)
			indices[entry.p2] = p2Index
		}
	}

	// Determine the people ordered by their id.
	people := make([]string, 0)
	for person := range indices {
		people = append(people, person)
	}
	sort.Slice(people, func(i, j int) bool {
		return indices[people[i]] < indices[people[j]]
	})

	// Now build the square matrix of weights
	weights := make([][]int, len(indices))
	for i := 0; i < len(indices); i++ {
		weights[i] = make([]int, len(indices))
	}
	for _, entry := range entries {
		weights[indices[entry.p1]][indices[entry.p2]] = entry.happiness
	}

	return people, weights
}

func Cost(perm []int, weights [][]int) int {
	N := len(perm)
	happiness := 0
	for i := 0; i < N; i++ {
		me := perm[i]
		left := perm[((i - 1 + N) % N)]
		right := perm[((i + 1) % N)]
		happiness += weights[me][left] + weights[me][right]
	}

	return happiness
}
