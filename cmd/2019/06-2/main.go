package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	parents := make(map[string]string)
	for _, orbit := range InputToOrbits(2019, 6) {
		parents[orbit.object] = orbit.center
	}

	path := func(n string) []string {
		var path []string
		for n != "" {
			path = append(path, n)
			n = parents[n]
		}
		return path[1 : len(path)-1]
	}

	// Determine the closest common ancestor between YOU and SAN.  That's the
	// vertex in the graph where we'll stop going inwards to COM and start going
	// outwards towards SAN.
	pathYOU := path("YOU")
	pathSAN := path("SAN")

	yi, si := len(pathYOU)-1, len(pathSAN)-1
	for pathYOU[yi] == pathSAN[si] {
		yi--
		si--
	}

	// We stopped when the nodes were different, so we have to go back one step to
	// be at the closest common ancestor.
	yi++
	si++

	// At this point we're pointing at the closest common ancestor.  The number of
	// nodes to the left of where we're pointing are the nodes we had to travel
	// through in order to get to the closest common ancestor.  So the distance
	// that we have to travel to get to santa is the sum of them.
	distance := yi + si
	fmt.Printf("distance: %d\n", distance)
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
