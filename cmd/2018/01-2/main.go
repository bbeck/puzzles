package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	deltas := aoc.InputToInts(2018, 1)

	seen := make(map[int]bool)
	var frequency int
outer:
	for {
		for _, delta := range deltas {
			frequency += delta
			if seen[frequency] {
				break outer
			}

			seen[frequency] = true
		}
	}

	fmt.Printf("frequency: %d\n", frequency)
}
