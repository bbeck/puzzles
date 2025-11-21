package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	target := in.ToGrid2D(func(x, y int, s string) string { return s })

	grid := NewGrid2D[string](34, 34)
	grid.Fill(".")

	id := func(g Grid2D[string]) string { return g.String() }
	_, cycle := FindCycleWithIdentity(grid, Next, id)

	var counts []int
	for _, g := range cycle {
		var count int
		if Matches(g, target) {
			var sums = make(map[string]int)
			g.ForEach(func(_ int, _ int, s string) { sums[s]++ })
			count = sums["#"]
		}

		counts = append(counts, count)
	}

	const N = 1000000000
	var div, mod = N / len(counts), N % len(counts)
	fmt.Println(div*Sum(counts...) + Sum(counts[:mod]...))
}

func Matches(grid Grid2D[string], target Grid2D[string]) bool {
	startX := (grid.Width - target.Width) / 2
	startY := (grid.Height - target.Height) / 2

	for x := startX; x < startX+target.Width; x++ {
		for y := startY; y < startY+target.Height; y++ {
			if grid.Get(x, y) != target.Get(x-startX, y-startY) {
				return false
			}
		}
	}
	return true
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
