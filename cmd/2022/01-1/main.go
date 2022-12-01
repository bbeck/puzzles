package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	calories := aoc.InputLinesTo(2022, 1, func(line string) (int, error) {
		if line == "" {
			return 0, nil
		}
		return aoc.ParseInt(line), nil
	})

	var best int
	for _, group := range aoc.Split(calories, func(n int) bool { return n != 0 }) {
		best = aoc.Max(best, aoc.Sum(group...))
	}
	fmt.Println(best)
}
