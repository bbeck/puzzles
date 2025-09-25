package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	numPriests := in.Int()
	numAcolytes := 1111
	blocks := 20240000
	blocks-- // Initially placed block

	width := 1
	for thickness := 1; blocks > 0; blocks -= thickness * width {
		width += 2
		thickness = (thickness * numPriests) % numAcolytes
	}

	fmt.Println(Abs(blocks) * width)
}
