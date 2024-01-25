package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	positions := InputToPositions()

	best := math.MaxInt
	for _, target := range positions {
		var cost int
		for _, p := range positions {
			cost += aoc.Abs(target - p)
		}
		best = aoc.Min(best, cost)
	}
	fmt.Println(best)
}

func InputToPositions() []int {
	line := aoc.InputToString(2021, 7)

	var fs []int
	for _, s := range strings.Split(line, ",") {
		fs = append(fs, aoc.ParseInt(s))
	}
	return fs
}
