package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

type Replacement struct {
	lhs, rhs string
}

func (r Replacement) Unapply(s string) string {
	index := strings.Index(s, r.rhs)
	if index != -1 {
		return strings.Replace(s, r.rhs, r.lhs, 1)
	}

	return s
}

func main() {
	replacements, molecule := InputToReplacementsAndMolecule(2015, 19)

	var count int
	for molecule != "e" {
		for _, replacement := range replacements {
			updated := replacement.Unapply(molecule)
			if updated != molecule {
				count++
				molecule = updated
			}
		}
	}

	fmt.Printf("count: %d\n", count)
}

func InputToReplacementsAndMolecule(year, day int) ([]Replacement, string) {
	lines := aoc.InputToLines(year, day)

	var replacements []Replacement
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			break
		}

		parts := strings.SplitN(lines[i], " => ", 2)
		lhs, rhs := parts[0], parts[1]

		replacements = append(replacements, Replacement{lhs, rhs})
	}

	molecule := lines[len(lines)-1]

	return replacements, molecule
}
