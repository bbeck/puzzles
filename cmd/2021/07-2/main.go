package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	positions := InputToPositions()

	choices := make(map[int]interface{})
	for _, position := range positions {
		choices[position] = nil
	}

	var best int
	for choice := range choices {
		var cost int
		for _, position := range positions {
			n := Abs(choice - position)
			cost += n * (n + 1) / 2
		}

		if best == 0 || cost < best {
			best = cost
		}
	}

	fmt.Println(best)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func InputToPositions() []int {
	line := aoc.InputToString(2021, 7)

	var fs []int
	for _, s := range strings.Split(strings.TrimSpace(line), ",") {
		fs = append(fs, aoc.ParseInt(s))
	}
	return fs
}
