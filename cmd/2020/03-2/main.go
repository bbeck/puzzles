package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := InputToGrid()
	slopes := []aoc.Point2D{
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

func CountTrees(grid aoc.Grid2D[bool], slope aoc.Point2D) int {
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

func InputToGrid() aoc.Grid2D[bool] {
	return aoc.InputToGrid2D(2020, 3, func(x int, y int, s string) bool {
		return s == "#"
	})
}
