package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"sort"
)

func main() {
	calories := puz.InputLinesTo(2022, 1, func(line string) int {
		if line == "" {
			return 0
		}
		return puz.ParseInt(line)
	})

	var groups []int
	for _, group := range puz.Split(calories, func(n int) bool { return n != 0 }) {
		groups = append(groups, puz.Sum(group...))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(groups)))
	fmt.Println(groups[0] + groups[1] + groups[2])
}
