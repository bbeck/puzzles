package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	stride := lib.InputToInt()

	var ring lib.Ring[int]
	for n := 0; n <= 2017; n++ {
		ring.NextN(stride)
		ring.InsertAfter(n)
	}

	ring.Next()
	fmt.Println(ring.Current())
}
