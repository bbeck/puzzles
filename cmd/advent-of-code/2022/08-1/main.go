package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	grid := lib.InputToIntGrid2D()

	var count int
	grid.ForEachPoint(func(p lib.Point2D, _ int) {
		visible := []bool{
			IsVisible(grid, p, lib.Left),
			IsVisible(grid, p, lib.Right),
			IsVisible(grid, p, lib.Up),
			IsVisible(grid, p, lib.Down),
		}
		if lib.Any(visible, lib.Identity[bool]) {
			count++
		}
	})
	fmt.Println(count)
}

func IsVisible(g lib.Grid2D[int], p lib.Point2D, h lib.Heading) bool {
	height := g.GetPoint(p)
	for p = p.Move(h); g.InBoundsPoint(p); p = p.Move(h) {
		if g.GetPoint(p) >= height {
			return false
		}
	}
	return true
}
