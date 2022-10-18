package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	graph := make(map[string]string)
	for _, o := range InputToOrbits() {
		graph[o.Child] = o.Parent
	}

	var count int
	for n := range graph {
		for {
			n = graph[n]
			if n == "" {
				break
			}
			count++
		}
	}
	fmt.Println(count)
}

type Orbit struct {
	Parent, Child string
}

func InputToOrbits() []Orbit {
	return aoc.InputLinesTo(2019, 6, func(line string) (Orbit, error) {
		var orbit Orbit
		orbit.Parent, orbit.Child, _ = strings.Cut(line, ")")
		return orbit, nil
	})
}
