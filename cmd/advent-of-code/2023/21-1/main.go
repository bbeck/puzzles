package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	grid, start := InputToGridAndStartingLocation()
	fmt.Println(Count(grid, start, 64))
}

func Count(g lib.Grid2D[string], p lib.Point2D, n int) int {
	current := lib.SetFrom(p)
	for n > 0 {
		var next lib.Set[lib.Point2D]
		for p := range current {
			g.ForEachOrthogonalNeighbor(p, func(q lib.Point2D, s string) {
				if s != "#" {
					next.Add(q)
				}
			})
		}

		current = next
		n--
	}

	return len(current)
}

func InputToGridAndStartingLocation() (lib.Grid2D[string], lib.Point2D) {
	grid := lib.InputToStringGrid2D()

	var start lib.Point2D
	grid.ForEachPoint(func(p lib.Point2D, s string) {
		if s == "S" {
			start = p
		}
	})

	return grid, start
}
