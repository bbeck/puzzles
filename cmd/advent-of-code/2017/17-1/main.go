package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	stride := puz.InputToInt()

	var ring puz.Ring[int]
	for n := 0; n <= 2017; n++ {
		ring.NextN(stride)
		ring.InsertAfter(n)
	}

	ring.Next()
	fmt.Println(ring.Current())
}
