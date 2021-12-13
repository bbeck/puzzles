package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	caves := InputToCaves()

	count := CountPaths(
		caves["start"],
		caves["end"],
		aoc.NewSingletonSet("start"),
		true,
	)
	fmt.Println(count)
}

func CountPaths(current *Cave, goal *Cave, seen aoc.Set, allow bool) int {
	if current == goal {
		return 1
	}

	var count int
	for _, neighbor := range current.Neighbors {
		if !neighbor.IsSmall {
			count += CountPaths(neighbor, goal, seen, allow)
		} else if !seen.Contains(neighbor.Name) {
			count += CountPaths(neighbor, goal, seen.Union(aoc.NewSingletonSet(neighbor.Name)), allow)
		} else if allow && neighbor.Name != "start" {
			count += CountPaths(neighbor, goal, seen, false)
		}
	}
	return count
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

		lhs.Neighbors = append(lhs.Neighbors, rhs)
		rhs.Neighbors = append(rhs.Neighbors, lhs)
	}

	return caves
}
