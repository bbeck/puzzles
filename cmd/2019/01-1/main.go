package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var fuel int
	for _, line := range aoc.InputToLines(2019, 1) {
		mass := aoc.ParseInt(line)
		fuel += mass/3 - 2
	}

	fmt.Printf("fuel: %d\n", fuel)
}
