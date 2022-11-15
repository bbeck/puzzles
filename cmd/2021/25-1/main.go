package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var step int
	for grid, changed := InputToGrid(), true; changed; step++ {
		grid, changed = Next(grid)
	}
	fmt.Println(step)
}

func Next(grid aoc.Grid2D[string]) (aoc.Grid2D[string], bool) {
	next := aoc.NewGrid2D[string](grid.Width, grid.Height)
	changed := false

	// Right
	grid.ForEach(func(x, y int, value string) {
		if value != ">" {
			if next.Get(x, y) == "" {
				next.Add(x, y, value)
			}
			return
		}

		nx := (x + 1) % grid.Width
		if grid.Get(nx, y) == "." {
			next.Add(x, y, ".")
			next.Add(nx, y, ">")
			changed = true
		} else {
			next.Add(x, y, ">")
		}
	})

	grid = next
	next = aoc.NewGrid2D[string](grid.Width, grid.Height)

	// Down
	grid.ForEach(func(x, y int, value string) {
		if value != "v" {
			if next.Get(x, y) == "" {
				next.Add(x, y, value)
			}
			return
		}

		ny := (y + 1) % grid.Height
		if grid.Get(x, ny) == "." {
			next.Add(x, y, ".")
			next.Add(x, ny, "v")
			changed = true
		} else {
			next.Add(x, y, "v")
		}
	})

	return next, changed
}

func InputToGrid() aoc.Grid2D[string] {
	return aoc.InputToGrid2D(2021, 25, func(x int, y int, s string) string {
		return s
	})
}
