package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	grid := InputToGrid()

	var count int
	for y := 0; y < grid.Height; y++ {
		x := 3 * y % grid.Width
		if grid.Get(x, y) {
			count++
		}
	}
	fmt.Println(count)
}

func InputToGrid() puz.Grid2D[bool] {
	return puz.InputToGrid2D(func(x int, y int, s string) bool {
		return s == "#"
	})
}
