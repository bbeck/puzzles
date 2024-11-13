package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	grid := puz.InputToIntGrid2D(2022, 8)

	var best int
	grid.ForEachPoint(func(p puz.Point2D, _ int) {
		counts := []int{
			Count(grid, p, puz.Left),
			Count(grid, p, puz.Right),
			Count(grid, p, puz.Up),
			Count(grid, p, puz.Down),
		}
		best = puz.Max(best, puz.Product(counts...))
	})
	fmt.Println(best)
}

func Count(g puz.Grid2D[int], p puz.Point2D, h puz.Heading) int {
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
