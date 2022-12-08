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
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			if IsVisible(grid, x, y) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func IsVisible(g aoc.Grid2D[int], x, y int) bool {
	visible := func(p aoc.Point2D, h aoc.Heading) bool {
		for p := p.Move(h); g.InBoundsPoint(p); p = p.Move(h) {
			if g.GetPoint(p) >= g.Get(x, y) {
				return false
			}
		}
		return true
	}

	p := aoc.Point2D{X: x, Y: y}
	return visible(p, aoc.Left) || visible(p, aoc.Right) || visible(p, aoc.Up) || visible(p, aoc.Down)
}
