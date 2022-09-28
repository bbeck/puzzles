package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	stride := aoc.InputToInt(2017, 17)

	var ring aoc.Ring[int]
	for n := 0; n <= 2017; n++ {
		ring.NextN(stride)
		ring.InsertAfter(n)
	}

	ring.Next()
	fmt.Println(ring.Current())
}
