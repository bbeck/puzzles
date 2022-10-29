package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, group := range InputToGroups(2020, 6) {
		intersection := group[0]
		for _, answers := range group {
			intersection = intersection.Intersect(answers)
		}

		count += len(intersection)
	}
	fmt.Println(count)
}

func InputToGroups(year, day int) [][]aoc.Set[string] {
	var groups [][]aoc.Set[string]

	var current []aoc.Set[string]
	for _, line := range aoc.InputToLines(year, day) {
		if len(line) == 0 {
			groups = append(groups, current)
			current = []aoc.Set[string]{}
			continue
		}

		// Each line is a single person's answers
		var answers aoc.Set[string]
		for _, question := range line {
			answers.Add(string(question))
		}
		current = append(current, answers)
	}
	groups = append(groups, current)

	return groups
}
