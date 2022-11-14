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
	for _, target := range aoc.SetFrom(positions...).Entries() {
		var cost int
		for _, p := range positions {
			n := aoc.Abs(target - p)
			cost += n * (n + 1) / 2
		}
		best = aoc.Min(best, cost)
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
