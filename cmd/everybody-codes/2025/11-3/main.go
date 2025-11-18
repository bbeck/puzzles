package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	ducks := in.Ints()

	// The key observation here is that we only need to simulate phase 2 since
	// the inputted numbers are in sorted order.  To simulate phase 2 we know
	// the results are going to end up with every column being equal.  We also
	// know that each round moves a number from one column to another, so we can
	// just count how many numbers need to be incremented to the mean.
	mean := Sum(ducks...) / len(ducks)

	var count int
	for _, c := range ducks {
		if c < mean {
			count += mean - c
		}
	}
	fmt.Println(count)
}
