package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	grid, start := InputToGridAndStartingLocation()
	fmt.Println(Count(grid, start, 64))
}

func Count(g puz.Grid2D[string], p puz.Point2D, n int) int {
	current := puz.SetFrom(p)
	for n > 0 {
		var next puz.Set[puz.Point2D]
		for p := range current {
			g.ForEachOrthogonalNeighbor(p, func(q puz.Point2D, s string) {
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

func InputToGridAndStartingLocation() (puz.Grid2D[string], puz.Point2D) {
	grid := puz.InputToStringGrid2D(2023, 21)

	var start puz.Point2D
	grid.ForEachPoint(func(p puz.Point2D, s string) {
		if s == "S" {
			start = p
		}
	})

	return grid, start
}
