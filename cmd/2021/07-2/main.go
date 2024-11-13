package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	positions := InputToPositions()

	best := math.MaxInt
	for _, target := range positions {
		var cost int
		for _, p := range positions {
			n := puz.Abs(target - p)
			cost += n * (n + 1) / 2
		}
		best = puz.Min(best, cost)
	}
	fmt.Println(best)
}

func InputToPositions() []int {
	line := puz.InputToString(2021, 7)

	var fs []int
	for _, s := range strings.Split(line, ",") {
		fs = append(fs, puz.ParseInt(s))
	}
	return fs
}
