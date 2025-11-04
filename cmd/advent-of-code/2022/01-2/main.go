package main

import (
	"fmt"
	"sort"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	calories := lib.InputLinesTo(func(line string) int {
		if line == "" {
			return 0
		}
		return lib.ParseInt(line)
	})

	var groups []int
	for _, group := range lib.Partition(calories, func(n int) bool { return n != 0 }) {
		groups = append(groups, lib.Sum(group...))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(groups)))
	fmt.Println(groups[0] + groups[1] + groups[2])
}
