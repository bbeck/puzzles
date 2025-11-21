package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToGrid2D(func(x, y int, s string) string { return s })

	var sums = make(map[string]int)
	for range 10 {
		grid = Next(grid)
		grid.ForEach(func(_ int, _ int, s string) { sums[s]++ })
	}
	fmt.Println(sums["#"])
}

func Next(g Grid2D[string]) Grid2D[string] {
	var TRANSITION = map[string]map[int]string{
		"#": {0: ".", 1: "#"},
		".": {0: "#", 1: "."},
	}

	next := NewGrid2D[string](g.Width, g.Height)
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			var current = g.Get(x, y)

			var values = make(map[string]int)
			if g.InBounds(x-1, y-1) {
				values[g.Get(x-1, y-1)]++
			}
			if g.InBounds(x-1, y+1) {
				values[g.Get(x-1, y+1)]++
			}
			if g.InBounds(x+1, y-1) {
				values[g.Get(x+1, y-1)]++
			}
			if g.InBounds(x+1, y+1) {
				values[g.Get(x+1, y+1)]++
			}
			next.Set(x, y, TRANSITION[current][values["#"]%2])
		}
	}
	return next
}
