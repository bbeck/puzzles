package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	caves := InputToCaves()

	paths := aoc.NewSet()
	FindPaths(
		caves["start"],
		caves["end"],
		&Path{
			names: []string{"start"},
			used:  map[string]int{"start": 1},
		},
		func(p *Path) {
			paths.Add(fmt.Sprintf("%s", p.names))
		},
	)

	fmt.Println(paths.Size())
}

func FindPaths(current *Cave, goal *Cave, path *Path, fn func(*Path)) {
	if current == goal {
		fn(path)
		return
	}

	for _, neighbor := range current.Neighbors {
		if path.used[neighbor.Name] > 0 && path.used["twice"] > 0 {
			continue
		}

		FindPaths(neighbor, goal, path.extend(neighbor), fn)
	}
}

type Path struct {
	names []string
	used  map[string]int
}

func (p *Path) extend(c *Cave) *Path {
	name := c.Name
	names := append(append([]string{}, p.names...), name)

	if !c.IsSmall {
		return &Path{
			names: names,
			used:  p.used,
		}
	}

	used := make(map[string]int)
	for n, c := range p.used {
		used[n] = c
		if n == name {
			used["twice"] = 1
		}
	}
	used[name] = 1

	return &Path{
		names: names,
		used:  used,
	}
}

type Cave struct {
	Name      string
	IsSmall   bool
	Neighbors []*Cave
}

func InputToCaves() map[string]*Cave {
	caves := make(map[string]*Cave)

	get := func(name string) *Cave {
		cave := caves[name]
		if cave == nil {
			cave = &Cave{
				Name:    name,
				IsSmall: strings.ToLower(name) == name,
			}
			caves[name] = cave
		}

		return cave
	}

	for _, line := range aoc.InputToLines(2021, 12) {
		parts := strings.Split(line, "-")
		lhs := get(parts[0])
		rhs := get(parts[1])

		if rhs.Name != "start" {
			lhs.Neighbors = append(lhs.Neighbors, rhs)
		}
		if lhs.Name != "start" {
			rhs.Neighbors = append(rhs.Neighbors, lhs)
		}
	}

	return caves
}
