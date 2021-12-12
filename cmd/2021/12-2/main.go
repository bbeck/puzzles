package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	caves := InputToCaves()

	paths := aoc.NewSet()
	Count(caves["start"], caves["end"], aoc.NewSet(), "", nil, paths)
	for name := range caves {
		if name == "start" || name == "end" || !IsLower(name) {
			continue
		}

		Count(caves["start"], caves["end"], aoc.NewSet(), name, nil, paths)
	}
	fmt.Println(paths.Size())
}

func Count(current, end *Cave, seen aoc.Set, twice string, path []string, paths aoc.Set) int {
	if current.Name != twice && seen.Contains(current.Name) {
		return 0
	}

	if current.Name == twice && seen.Contains("used") {
		return 0
	}

	if current.Name == end.Name {
		s := fmt.Sprintf("%v", path)
		paths.Add(s)
		return 1
	}

	if IsLower(current.Name) && !seen.Add(current.Name) {
		seen.Add("used")
	}

	var ways int
	for _, neighbor := range current.Neighbors {
		ways += Count(neighbor, end, seen.Union(aoc.NewSet()), twice, append(path, neighbor.Name), paths)
	}
	return ways
}

func IsLower(s string) bool {
	return s[0] >= 'a' && s[0] <= 'z'
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
