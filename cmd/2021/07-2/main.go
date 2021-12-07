package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	positions := InputToPositions()

	choices := make(map[int]interface{})
	for _, position := range positions {
		choices[position] = nil
	}

	best := math.MaxInt
	for choice := range choices {
		var cost int
		for _, position := range positions {
			n := aoc.AbsInt(choice - position)
			cost += n * (n + 1) / 2
		}

		best = aoc.MinInt(best, cost)
	}

	fmt.Println(best)
}

func InputToPositions() []int {
	line := aoc.InputToString(2021, 7)

	var fs []int
	for _, s := range strings.Split(strings.TrimSpace(line), ",") {
		fs = append(fs, aoc.ParseInt(s))
	}
	return fs
}
