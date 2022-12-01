package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var best, group int
	for _, line := range aoc.InputToLines(2022, 1) {
		if line != "" {
			group += aoc.ParseInt(line)
			continue
		}

		best = aoc.Max(best, group)
		group = 0
	}
	fmt.Println(best)
}
