package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	calories := aoc.InputLinesTo(2022, 1, func(line string) int {
		if line == "" {
			return 0
		}
		return aoc.ParseInt(line)
	})

	var best int
	for _, group := range aoc.Split(calories, func(n int) bool { return n != 0 }) {
		best = aoc.Max(best, aoc.Sum(group...))
	}
	fmt.Println(best)
}
