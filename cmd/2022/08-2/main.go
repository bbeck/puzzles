package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := aoc.InputToGrid2D[int](2022, 8, func(x, y int, value string) int {
		return aoc.ParseInt(value)
	})

	var score int
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			score = aoc.Max(score, Score(grid, x, y))
		}
	}
	fmt.Println(score)
}

func Score(g aoc.Grid2D[int], x, y int) int {
	count := func(p aoc.Point2D, h aoc.Heading) int {
		var n int
		for p = p.Move(h); g.InBoundsPoint(p); p = p.Move(h) {
			n++
			if g.GetPoint(p) >= g.Get(x, y) {
				break
			}
		}
		return n
	}

	p := aoc.Point2D{X: x, Y: y}
	return count(p, aoc.Left) * count(p, aoc.Right) * count(p, aoc.Up) * count(p, aoc.Down)
}
