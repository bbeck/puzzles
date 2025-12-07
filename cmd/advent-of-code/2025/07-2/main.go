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

var memo = make(map[Point2D]int)

func Count(grid Grid2D[string], p Point2D) int {
	if v, ok := memo[p]; ok {
		return v
	}

	if p.Y == grid.Height-1 {
		return 1
	}

	var v int
	if grid.InBoundsPoint(p) && grid.GetPoint(p) == "^" {
		v = Count(grid, p.Left()) + Count(grid, p.Right())
	} else {
		v = Count(grid, p.Down())
	}

	memo[p] = v
	return v
}
