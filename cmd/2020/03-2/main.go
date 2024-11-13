package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	grid := InputToGrid()
	slopes := []puz.Point2D{
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

func CountTrees(grid puz.Grid2D[bool], slope puz.Point2D) int {
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

func InputToGrid() puz.Grid2D[bool] {
	return puz.InputToGrid2D(2020, 3, func(x int, y int, s string) bool {
		return s == "#"
	})
}
