package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"sort"
)

func main() {
	var groups []int
	var group int
	for _, line := range aoc.InputToLines(2022, 1) {
		if line != "" {
			group += aoc.ParseInt(line)
			continue
		}

		groups = append(groups, group)
		group = 0
	}

	sort.Sort(sort.Reverse(sort.IntSlice(groups)))
	fmt.Println(groups[0] + groups[1] + groups[2])
}
