package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/lib"
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

func InputToGroups(year, day int) [][]lib.Set[string] {
	var groups [][]lib.Set[string]

	var current []lib.Set[string]
	for _, line := range lib.InputToLines() {
		if len(line) == 0 {
			groups = append(groups, current)
			current = []lib.Set[string]{}
			continue
		}

		// Each line is a single person's answers
		var answers lib.Set[string]
		for _, question := range line {
			answers.Add(string(question))
		}
		current = append(current, answers)
	}
	groups = append(groups, current)

	return groups
}
