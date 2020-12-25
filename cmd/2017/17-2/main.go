package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	stride := aoc.InputToInt(2017, 17)

	// Remember the value after the 0
	var after int

	index := 1
	for n := 1; n <= 50_000_000; n++ {
		// Calculate the index to insert at.
		index = (index+stride)%n + 1

		// If we're inserting into index 1, then we're inserting right after the 0,
		// remember the value being inserted.
		if index == 1 {
			after = n
		}
	}

	fmt.Printf("value: %d\n", after)
}
