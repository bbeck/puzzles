package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	blocks := in.Int()
	blocks-- // Initially placed block

	var width int
	for width = 1; blocks > 0; blocks -= width {
		width += 2
	}

	fmt.Println(width * Abs(blocks))
}
