package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

type Replacement struct {
	lhs, rhs string
}

func (r Replacement) Apply(s string) []string {
	var replaced []string
	for _, i := range Indices(s, r.lhs) {
		replaced = append(replaced, s[0:i]+r.rhs+s[i+len(r.lhs):])
	}
	return replaced
}

func main() {
	replacements, molecule := InputToReplacementsAndMolecule(2015, 19)

	molecules := make(map[string]int)
	for _, replacement := range replacements {
		for _, newMolecule := range replacement.Apply(molecule) {
			molecules[newMolecule]++
		}
	}

	fmt.Printf("num distinct molecules: %d\n", len(molecules))
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

func Indices(s string, substr string) []int {
	indices := make([]int, 0)

	offset := 0
	for {
		index := strings.Index(s[offset:], substr)
		if index == -1 {
			break
		}

		indices = append(indices, offset+index)
		offset += index + 1
	}

	return indices
}
