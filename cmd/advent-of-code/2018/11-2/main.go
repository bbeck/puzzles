package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	serial := lib.InputToInt()

	// Partial sums that sum all cells above and to the left of the current one
	grid := lib.NewGrid2D[int](301, 301)
	for x := 1; x < grid.Width; x++ {
		rack := x + 10

		for y := 1; y < grid.Height; y++ {
			cell := (rack*y + serial) * rack
			cell = (cell%1000)/100 - 5

			sum := cell + grid.Get(x-1, y) + grid.Get(x, y-1) - grid.Get(x-1, y-1)
			grid.Set(x, y, sum)
		}
	}

	var best int
	var p lib.Point2D
	var size int
	for n := 1; n <= grid.Width; n++ {
		for x := 1; x < grid.Width-n; x++ {
			for y := 1; y < grid.Height-n; y++ {
				total := grid.Get(x+n-1, y+n-1) - grid.Get(x-1, y+n-1) - grid.Get(x+n-1, y-1) + grid.Get(x-1, y-1)
				if total > best {
					best = total
					p = lib.Point2D{X: x, Y: y}
					size = n
				}
			}
		}
	}

	fmt.Printf("%d,%d,%d\n", p.X, p.Y, size)
}
