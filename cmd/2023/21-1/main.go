package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid, start := InputToGridAndStartingLocation()
	fmt.Println(Count(grid, start, 64))
}

func Count(g aoc.Grid2D[string], p aoc.Point2D, n int) int {
	current := aoc.SetFrom(p)
	for n > 0 {
		var next aoc.Set[aoc.Point2D]
		for p := range current {
			g.ForEachOrthogonalNeighbor(p, func(q aoc.Point2D, s string) {
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

func InputToGridAndStartingLocation() (aoc.Grid2D[string], aoc.Point2D) {
	grid := aoc.InputToStringGrid2D(2023, 21)

	var start aoc.Point2D
	grid.ForEachPoint(func(p aoc.Point2D, s string) {
		if s == "S" {
			start = p
		}
	})

	return grid, start
}
