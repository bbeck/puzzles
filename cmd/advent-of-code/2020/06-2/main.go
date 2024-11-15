package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
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

func InputToGroups(year, day int) [][]puz.Set[string] {
	var groups [][]puz.Set[string]

	var current []puz.Set[string]
	for _, line := range puz.InputToLines() {
		if len(line) == 0 {
			groups = append(groups, current)
			current = []puz.Set[string]{}
			continue
		}

		// Each line is a single person's answers
		var answers puz.Set[string]
		for _, question := range line {
			answers.Add(string(question))
		}
		current = append(current, answers)
	}
	groups = append(groups, current)

	return groups
}
