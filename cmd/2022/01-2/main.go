package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"sort"
)

func main() {
	calories := aoc.InputLinesTo(2022, 1, func(line string) (int, error) {
		if line == "" {
			return 0, nil
		}
		return aoc.ParseInt(line), nil
	})

	var groups []int
	for _, group := range aoc.Split(calories, func(n int) bool { return n != 0 }) {
		groups = append(groups, aoc.Sum(group...))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(groups)))
	fmt.Println(groups[0] + groups[1] + groups[2])
}
