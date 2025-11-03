package main

import (
	"fmt"
	"math"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	// Looking at the grids there appears to always be at least one column that
	// contains nothing but warm air streams.  The goal is going to be to get to
	// the closest of those columns to the start as quickly as possible and then
	// exploit it for as long as possible.
	var grid = in.ToGrid2D(func(_, _ int, s string) string { return s })
	var W = grid.Width
	var H = grid.Height
	var dA = map[string]int{".": +1, "S": +1, "+": -1, "-": +2, "#": +H}

	var start int
	grid.ForEach(func(x int, _ int, s string) {
		if s == "S" {
			start = x
		}
	})

	// Find the cheapest column to go down that is closest to the start.
	var col int
	var best = math.MaxInt
	for x := range W {
		var cost int
		for y := range H {
			cost += dA[grid.Get(x, y)]
		}

		if cost < best || (cost == best && Abs(start-x) < Abs(start-col)) {
			col = x
			best = cost
		}
	}

	// Account for the cost to get to our starting column.
	var altitude = 384400 - Abs(start-col)

	var y int
	for altitude > 0 {
		y++
		altitude -= dA[grid.Get(col, y%H)]
	}
	fmt.Println(y)
}
