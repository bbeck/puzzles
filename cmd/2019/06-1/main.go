package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	// We'll view this as a tree and then count the number of nodes in a traversal
	// from a child to the root.
	parents := make(map[string]string)
	for _, orbit := range InputToOrbits(2019, 6) {
		parents[orbit.object] = orbit.center
	}

	count := func(n string) int {
		var count int
		for n != "" {
			n = parents[n]
			count++
		}
		return count - 1
	}

	var sum int
	for node := range parents {
		sum += count(node)
	}

	fmt.Printf("total number of orbits: %d\n", sum)
}

type Orbit struct {
	object string
	center string
}

func InputToOrbits(year, day int) []Orbit {
	var orbits []Orbit
	for _, line := range aoc.InputToLines(year, day) {
		parts := strings.Split(line, ")")
		center, object := parts[0], parts[1]

		orbits = append(orbits, Orbit{object: object, center: center})
	}

	return orbits
}
