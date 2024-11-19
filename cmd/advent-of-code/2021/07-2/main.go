package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	positions := InputToPositions()

	best := math.MaxInt
	for _, target := range positions {
		var cost int
		for _, p := range positions {
			n := lib.Abs(target - p)
			cost += n * (n + 1) / 2
		}
		best = lib.Min(best, cost)
	}
	fmt.Println(best)
}

func InputToPositions() []int {
	line := lib.InputToString()

	var fs []int
	for _, s := range strings.Split(line, ",") {
		fs = append(fs, lib.ParseInt(s))
	}
	return fs
}
