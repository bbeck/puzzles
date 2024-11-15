package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
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
	return puz.InputLinesTo(func(line string) Orbit {
		var orbit Orbit
		orbit.Parent, orbit.Child, _ = strings.Cut(line, ")")
		return orbit
	})
}
