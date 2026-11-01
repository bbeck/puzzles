package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var grid = in.ToStringGrid2D()

	// Land: grow from the center
	var center = Point2D{X: grid.Width / 2, Y: grid.Height / 2}
	grid.SetPoint(center, "#")
	Grow(grid, []Point2D{center}, "#", func(p Point2D, s string) bool {
		return s != "#" && Count(grid, p, "1", "#") >= 5
	})

	var count int
	grid.ForEach(func(_ int, _ int, s string) {
		if s == "#" {
			count++
		}
	})
	fmt.Println(count)
}

func Grow(grid Grid2D[string], start []Point2D, symbol string, fn func(Point2D, string) bool) {
	var queue = start
	for len(queue) > 0 {
		var p Point2D
		p, queue = queue[0], queue[1:]

		for _, q := range p.Neighbors() {
			if grid.InBoundsPoint(q) && fn(q, grid.GetPoint(q)) {
				grid.SetPoint(q, symbol)
				queue = append(queue, q)
			}
		}
	}
}

func Count(grid Grid2D[string], p Point2D, needles ...string) int {
	var set = SetFrom(needles...)

	var count int
	grid.ForEachNeighborPoint(p, func(_ Point2D, s string) {
		if set.Contains(s) {
			count++
		}
	})
	return count
}
