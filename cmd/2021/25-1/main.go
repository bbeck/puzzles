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
	grid.ForEachXY(func(x, y int, value string) {
		if value != ">" {
			if next.GetXY(x, y) == "" {
				next.AddXY(x, y, value)
			}
			return
		}

		nx := (x + 1) % grid.Width
		if grid.GetXY(nx, y) == "." {
			next.AddXY(x, y, ".")
			next.AddXY(nx, y, ">")
			changed = true
		} else {
			next.AddXY(x, y, ">")
		}
	})

	grid = next
	next = aoc.NewGrid2D[string](grid.Width, grid.Height)

	// Down
	grid.ForEachXY(func(x, y int, value string) {
		if value != "v" {
			if next.GetXY(x, y) == "" {
				next.AddXY(x, y, value)
			}
			return
		}

		ny := (y + 1) % grid.Height
		if grid.GetXY(x, ny) == "." {
			next.AddXY(x, y, ".")
			next.AddXY(x, ny, "v")
			changed = true
		} else {
			next.AddXY(x, y, "v")
		}
	})

	return next, changed
}

func InputToGrid() aoc.Grid2D[string] {
	lines := aoc.InputToLines(2021, 25)

	grid := aoc.NewGrid2D[string](len(lines[0]), len(lines))
	for y, line := range lines {
		for x, c := range line {
			grid.AddXY(x, y, string(c))
		}
	}

	return grid
}
