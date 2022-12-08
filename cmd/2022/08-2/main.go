package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := aoc.InputToGrid2D[int](2022, 8, func(x, y int, value string) int {
		return aoc.ParseInt(value)
	})

	var best int
	grid.ForEachPoint(func(p aoc.Point2D, _ int) {
		counts := []int{
			Count(grid, p, aoc.Left),
			Count(grid, p, aoc.Right),
			Count(grid, p, aoc.Up),
			Count(grid, p, aoc.Down),
		}
		best = aoc.Max(best, aoc.Product(counts...))
	})
	fmt.Println(best)
}

func Count(g aoc.Grid2D[int], p aoc.Point2D, h aoc.Heading) int {
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
