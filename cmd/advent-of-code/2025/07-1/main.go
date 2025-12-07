package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToStringGrid2D()

	var start Point2D
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "S" {
			start = p
		}
	})
	fmt.Println(Count(grid, start))
}

var seen Set[Point2D]

func Count(grid Grid2D[string], p Point2D) int {
	if !seen.Add(p) || !grid.InBoundsPoint(p) {
		return 0
	}

	if grid.GetPoint(p) != "^" {
		return Count(grid, p.Down())
	}

	return Count(grid, p.Left()) + Count(grid, p.Right()) + 1
}
