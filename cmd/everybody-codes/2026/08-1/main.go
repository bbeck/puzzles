package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

const MISSING = -1

func main() {
	grid := in.ToGrid2D(func(x, y int, s string) int {
		if s == "." {
			return MISSING
		}
		return ParseInt(s)
	})

	var count int
	grid.ForEachPoint(func(p Point2D, value int) {
		switch {
		case value != MISSING && InShadow(grid, p, value):
			grid.SetPoint(p, MISSING)
		case value != MISSING:
			count++
		}
	})
	fmt.Println(count)
}

func InShadow(grid Grid2D[int], p Point2D, height int) bool {
	for q, dist := p.Up().Left(), 1; grid.InBoundsPoint(q); q, dist = q.Up().Left(), dist+1 {
		other := grid.GetPoint(q)
		if other != MISSING && other-dist >= height {
			return true
		}
	}
	return false
}
