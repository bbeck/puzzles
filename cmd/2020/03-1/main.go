package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := InputToGrid()

	var count int
	for y := 0; y < grid.Height; y++ {
		x := 3 * y % grid.Width
		if grid.GetXY(x, y) {
			count++
		}
	}
	fmt.Println(count)
}

func InputToGrid() aoc.Grid2D[bool] {
	return aoc.InputToGrid2D(2020, 3, func(x int, y int, s string) bool {
		return s == "#"
	})
}
