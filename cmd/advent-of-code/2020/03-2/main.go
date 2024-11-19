package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	grid := InputToGrid()
	slopes := []lib.Point2D{
		{X: 1, Y: 1},
		{X: 3, Y: 1},
		{X: 5, Y: 1},
		{X: 7, Y: 1},
		{X: 1, Y: 2},
	}

	product := 1
	for _, slope := range slopes {
		product *= CountTrees(grid, slope)
	}
	fmt.Println(product)
}

func CountTrees(grid lib.Grid2D[bool], slope lib.Point2D) int {
	var count int

	var x, y int
	for y < grid.Height {
		if grid.Get(x, y) {
			count++
		}
		x = (x + slope.X) % grid.Width
		y += slope.Y
	}

	return count
}

func InputToGrid() lib.Grid2D[bool] {
	return lib.InputToGrid2D(func(x int, y int, s string) bool {
		return s == "#"
	})
}
