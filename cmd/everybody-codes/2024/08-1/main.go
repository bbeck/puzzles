package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	blocks := InputToInt()
	blocks-- // Initially placed block

	var width int
	for width = 1; blocks > 0; blocks -= width {
		width += 2
	}

	fmt.Println(width * Abs(blocks))
}
