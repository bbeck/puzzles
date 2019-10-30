package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	stride := aoc.InputToInt(2017, 17)

	ring := aoc.NewRing()
	ring.InsertAfter(0)

	for n := 1; n <= 2017; n++ {
		ring.NextN(stride)
		ring.InsertAfter(n)
	}

	ring.Next()
	fmt.Printf("current: %+v\n", ring.Current())
}
