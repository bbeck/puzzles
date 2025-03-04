package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib/in"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	stride := in.Int()

	var ring Ring[int]
	for n := 0; n <= 2017; n++ {
		ring.NextN(stride)
		ring.InsertAfter(n)
	}

	ring.Next()
	fmt.Println(ring.Current())
}
