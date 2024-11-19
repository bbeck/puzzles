package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	grid := lib.InputToIntGrid2D()

	var best int
	grid.ForEachPoint(func(p lib.Point2D, _ int) {
		counts := []int{
			Count(grid, p, lib.Left),
			Count(grid, p, lib.Right),
			Count(grid, p, lib.Up),
			Count(grid, p, lib.Down),
		}
		best = lib.Max(best, lib.Product(counts...))
	})
	fmt.Println(best)
}

func Count(g lib.Grid2D[int], p lib.Point2D, h lib.Heading) int {
	height := g.GetPoint(p)

	var count int
	for p = p.Move(h); g.InBoundsPoint(p); p = p.Move(h) {
		count++
		if g.GetPoint(p) >= height {
			break
		}
	}
	return count
}
