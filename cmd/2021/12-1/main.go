package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	caves := InputToCaves()
	count := Count(caves["start"], caves["end"], aoc.NewSet())
	fmt.Println(count)
}

func Count(current, end *Cave, seen aoc.Set) int {
	if seen.Contains(current.Name) {
		return 0
	}

	if current.Name == end.Name {
		return 1
	}

	if current.Name[0] >= 'a' && current.Name[0] <= 'z' {
		seen.Add(current.Name)
	}

	var ways int
	for _, neighbor := range current.Neighbors {
		ways += Count(neighbor, end, seen.Union(aoc.NewSet()))
	}
	return ways
}

type Cave struct {
	Name      string
	Neighbors []*Cave
}

func InputToCaves() map[string]*Cave {
	caves := make(map[string]*Cave)
	for _, line := range aoc.InputToLines(2021, 12) {
		parts := strings.Split(line, "-")
		lhs := caves[parts[0]]
		if lhs == nil {
			lhs = &Cave{Name: parts[0]}
			caves[parts[0]] = lhs
		}

		rhs := caves[parts[1]]
		if rhs == nil {
			rhs = &Cave{Name: parts[1]}
			caves[parts[1]] = rhs
		}

		lhs.Neighbors = append(lhs.Neighbors, rhs)
		rhs.Neighbors = append(rhs.Neighbors, lhs)
	}

	return caves
}
