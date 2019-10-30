package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	offsets := aoc.InputToInts(2017, 5)

	var steps int
	for pc := 0; pc >= 0 && pc < len(offsets); steps++ {
		offset := offsets[pc]
		if offset >= 3 {
			offsets[pc]--
		} else {
			offsets[pc]++
		}
		pc += offset
	}

	fmt.Printf("steps: %d\n", steps)
}
