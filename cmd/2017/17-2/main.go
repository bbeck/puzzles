package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	stride := aoc.InputToInt(2017, 17)

	ring := aoc.NewRing()
	ring.InsertAfter(0)

	for n := 1; n <= 50_000_000; n++ {
		ring.NextN(stride)
		ring.InsertAfter(n)
	}

	// Find the element after 0
	for current := ring.Current(); current.(int) != 0; current = ring.Next() {
	}
	fmt.Printf("value: %+v\n", ring.Next())
}
