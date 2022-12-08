package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := aoc.InputToGrid2D[int](2022, 8, func(x, y int, value string) int {
		return aoc.ParseInt(value)
	})

	var count int
	grid.ForEachPoint(func(p aoc.Point2D, _ int) {
		visible := []bool{
			IsVisible(grid, p, aoc.Left),
			IsVisible(grid, p, aoc.Right),
			IsVisible(grid, p, aoc.Up),
			IsVisible(grid, p, aoc.Down),
		}
		if aoc.Any(visible, aoc.Identity[bool]) {
			count++
		}
	})
	fmt.Println(count)
}

func IsVisible(g aoc.Grid2D[int], p aoc.Point2D, h aoc.Heading) bool {
	height := g.GetPoint(p)
	for p = p.Move(h); g.InBoundsPoint(p); p = p.Move(h) {
		if g.GetPoint(p) >= height {
			return false
		}
	}
	return true
}
