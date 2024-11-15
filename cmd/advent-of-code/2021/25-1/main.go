package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	grid := puz.InputToStringGrid2D()

	var step int
	for changed := true; changed; step++ {
		grid, changed = Next(grid)
	}
	fmt.Println(step)
}

func Next(grid puz.Grid2D[string]) (puz.Grid2D[string], bool) {
	next := puz.NewGrid2D[string](grid.Width, grid.Height)
	changed := false

	// Right
	grid.ForEach(func(x, y int, value string) {
		if value != ">" {
			if next.Get(x, y) == "" {
				next.Set(x, y, value)
			}
			return
		}

		nx := (x + 1) % grid.Width
		if grid.Get(nx, y) == "." {
			next.Set(x, y, ".")
			next.Set(nx, y, ">")
			changed = true
		} else {
			next.Set(x, y, ">")
		}
	})

	grid = next
	next = puz.NewGrid2D[string](grid.Width, grid.Height)

	// Down
	grid.ForEach(func(x, y int, value string) {
		if value != "v" {
			if next.Get(x, y) == "" {
				next.Set(x, y, value)
			}
			return
		}

		ny := (y + 1) % grid.Height
		if grid.Get(x, ny) == "." {
			next.Set(x, y, ".")
			next.Set(x, ny, "v")
			changed = true
		} else {
			next.Set(x, y, "v")
		}
	})

	return next, changed
}
