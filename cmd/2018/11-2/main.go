package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	serial := aoc.InputToInt(2018, 11)

	// Partial sums that sum all cells above and to the left of the current one
	grid := aoc.NewGrid2D[int](301, 301)
	for x := 1; x < grid.Width; x++ {
		rack := x + 10

		for y := 1; y < grid.Height; y++ {
			cell := (rack*y + serial) * rack
			cell = (cell%1000)/100 - 5

			sum := cell + grid.GetXY(x-1, y) + grid.GetXY(x, y-1) - grid.GetXY(x-1, y-1)
			grid.AddXY(x, y, sum)
		}
	}

	var best int
	var p aoc.Point2D
	var size int
	for n := 1; n <= grid.Width; n++ {
		for x := 1; x < grid.Width-n; x++ {
			for y := 1; y < grid.Height-n; y++ {
				total := grid.GetXY(x+n-1, y+n-1) - grid.GetXY(x-1, y+n-1) - grid.GetXY(x+n-1, y-1) + grid.GetXY(x-1, y-1)
				if total > best {
					best = total
					p = aoc.Point2D{X: x, Y: y}
					size = n
				}
			}
		}
	}

	fmt.Printf("%d,%d,%d\n", p.X, p.Y, size)
}
